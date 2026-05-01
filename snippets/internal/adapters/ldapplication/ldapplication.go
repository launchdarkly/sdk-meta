package ldapplication

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/atomicfile"
	"github.com/launchdarkly/sdk-meta/snippets/internal/markers"
	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
	"github.com/launchdarkly/sdk-meta/snippets/internal/version"
)

// markerSentinel is the substring every render marker contains. We use it as
// a fast pre-filter so walking large consumer trees only invokes the marker
// scanner on files that could possibly contain a marker. Cheaper than a
// regex match per file.
var markerSentinel = []byte("SDK_SNIPPET:RENDER:")

// candidateExtensions are the file extensions whose comment syntax the
// marker scanner understands today: JS, TS, JSX, TSX, plus MDX (whose JSX
// expression comments share the `{/* ... */}` form). Everything else is
// skipped on sight so a `node_modules` walk doesn't eat a million PNGs.
var candidateExtensions = map[string]struct{}{
	".tsx": {}, ".jsx": {}, ".ts": {}, ".js": {}, ".mdx": {},
}

// skipDirNames are directories we never descend into. They never carry
// snippet markers and routinely have hundreds of thousands of files
// (node_modules), generated build output, or version-control bookkeeping.
var skipDirNames = map[string]struct{}{
	"node_modules": {}, ".git": {}, ".next": {}, "dist": {}, "build": {},
	".cache": {}, "coverage": {}, "out": {}, ".turbo": {},
}

// Render walks every entrypoint, finds files that contain render markers,
// and rewrites each marked region with the rendered snippet content. sdksFS
// is the fs.FS rooted at the sdks/ directory (either embedded or
// os.DirFS(path)). entrypoints are absolute or repo-relative directory
// paths the consumer (gonfalon, ld-docs) declares as roots to scan.
// Returns one entry per file it touched.
func Render(sdksFS fs.FS, entrypoints []string) ([]string, error) {
	return run(sdksFS, entrypoints, false)
}

// Verify re-renders every marked region in memory and fails if any hash in a
// marker does not match the hash of the bytes currently in the file, or if a
// re-render would change content. Never modifies files.
func Verify(sdksFS fs.FS, entrypoints []string) error {
	_, err := run(sdksFS, entrypoints, true)
	return err
}

func run(sdksFS fs.FS, entrypoints []string, dryRun bool) ([]string, error) {
	snippets, err := model.LoadAll(sdksFS)
	if err != nil {
		return nil, err
	}

	files, err := discoverFilesUnder(entrypoints)
	if err != nil {
		return nil, err
	}

	var changed []string
	for _, p := range files {
		ok, err := rewriteFile(p, snippets, dryRun)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", p, err)
		}
		if ok {
			changed = append(changed, p)
		}
	}
	return changed, nil
}

// discoverFilesUnder walks every entrypoint directory and returns the set
// of files whose extension the marker scanner understands AND whose contents
// contain the SDK_SNIPPET:RENDER: sentinel. Symlinks are not followed so
// a malicious symlink farm in node_modules can't pull the renderer outside
// its intended scope. Duplicates (when two entrypoints overlap) are
// collapsed.
func discoverFilesUnder(entrypoints []string) ([]string, error) {
	if len(entrypoints) == 0 {
		return nil, fmt.Errorf("at least one --entrypoint is required")
	}
	seen := map[string]struct{}{}
	var out []string
	for _, ep := range entrypoints {
		abs, err := filepath.Abs(ep)
		if err != nil {
			return nil, fmt.Errorf("entrypoint %q: %w", ep, err)
		}
		info, err := os.Stat(abs)
		if err != nil {
			return nil, fmt.Errorf("entrypoint %q: %w", ep, err)
		}
		if !info.IsDir() {
			return nil, fmt.Errorf("entrypoint %q: not a directory", ep)
		}
		err = filepath.WalkDir(abs, func(p string, d os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return walkErr
			}
			if d.IsDir() {
				// Always descend into the entrypoint root, even if its
				// basename happens to be in skipDirNames (e.g.
				// `--entrypoint=./build` for a project that lays its
				// generated TSX out there). The skip-list is meant to prune
				// well-known noise *under* the root, not to silently turn
				// the entire walk into a no-op.
				if p != abs {
					if _, skip := skipDirNames[d.Name()]; skip {
						return filepath.SkipDir
					}
				}
				return nil
			}
			// Skip symlinks — we don't follow links into other parts of the
			// repo (or out of it).
			if d.Type()&os.ModeSymlink != 0 {
				return nil
			}
			if _, ok := candidateExtensions[filepath.Ext(p)]; !ok {
				return nil
			}
			if _, dup := seen[p]; dup {
				return nil
			}
			data, err := os.ReadFile(p)
			if err != nil {
				return err
			}
			if !bytes.Contains(data, markerSentinel) {
				return nil
			}
			seen[p] = struct{}{}
			out = append(out, p)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

// rewriteFile does the actual in-place rewrite. If dryRun is true it only
// verifies that (a) every marker carries a hash field, (b) that hash matches
// the full <Tag>...</Tag> region in the file, and (c) re-rendering would
// produce the same bytes.
func rewriteFile(path string, snippets map[string]*model.Snippet, dryRun bool) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	src := string(raw)

	matches, err := markers.ScanTSX(src)
	if err != nil {
		return false, err
	}
	if len(matches) == 0 {
		return false, nil
	}

	// Build output by replacing each match's region in left-to-right order.
	var sb strings.Builder
	cursor := 0
	changed := false
	for _, m := range matches {
		s := snippets[m.Fields.ID]
		if s == nil {
			return false, fmt.Errorf("marker %q references unknown snippet id", m.Fields.ID)
		}
		tpl, err := render.Parse(s.CodeBody)
		if err != nil {
			return false, fmt.Errorf("snippet %s: %w", s.Path, err)
		}
		declared := declaredInputSet(s)
		// Reuse the surrounding whitespace from the existing region so a
		// re-render produces a minimal diff. This is purely cosmetic; the
		// bare-vs-template decision below is independent and is driven by
		// the snippet's intent.
		leading, trailing := splitSurroundingWS(src[m.RegionStart:m.RegionEnd])
		jsxBody := leading + renderForJSXChild(tpl, declared) + trailing

		// The hash covers ONLY the children we own. Attributes on the
		// element are the consumer's to choose (lang="…", withCopyButton,
		// className, etc.) and re-styling them must not require re-running
		// `snippets render`. This matches the scope=content contract.
		newHash := markers.HashContent(jsxBody)

		if dryRun {
			if m.Fields.Hash == "" {
				return false, fmt.Errorf("marker %q: missing required hash= field — re-render to populate it", m.Fields.ID)
			}
			actualHash := markers.HashContent(src[m.RegionStart:m.RegionEnd])
			if m.Fields.Hash != actualHash {
				return false, fmt.Errorf("marker %q: hand-edit detected — children hash %s does not match marker %s",
					m.Fields.ID, actualHash, m.Fields.Hash)
			}
			if jsxBody != src[m.RegionStart:m.RegionEnd] {
				return false, fmt.Errorf("marker %q: re-render would change region — run `snippets render`", m.Fields.ID)
			}
			continue
		}

		// `version=` records the binary that last *changed* this snippet's
		// rendered content — not the binary that last touched the file. If
		// the body is byte-identical to what's already on disk, we preserve
		// the existing version so a release-without-content-changes doesn't
		// rewrite every marker (and produce a noisy 24-file diff in the
		// downstream sync PR). First-render markers with no `version=` field
		// get stamped with the current binary's version.
		ver := m.Fields.Version
		if ver == "" || jsxBody != src[m.RegionStart:m.RegionEnd] {
			ver = version.Version
		}
		newMarker := markers.FormatMarker(m.Style, markers.MarkerFields{
			ID:      m.Fields.ID,
			Hash:    newHash,
			Version: ver,
			Scope:   "content",
		})
		sb.WriteString(src[cursor:m.CommentStart])
		sb.WriteString(newMarker)
		sb.WriteString(src[m.CommentEnd:m.RegionStart])
		sb.WriteString(jsxBody)
		cursor = m.RegionEnd

		if src[m.CommentStart:m.CommentEnd] != newMarker || src[m.RegionStart:m.RegionEnd] != jsxBody {
			changed = true
		}
	}
	if dryRun {
		return false, nil
	}
	sb.WriteString(src[cursor:])
	if !changed {
		return false, nil
	}
	return true, atomicfile.Write(path, []byte(sb.String()))
}

// renderForJSXChild produces the bytes that go between <Tag> and </Tag>.
// The choice between bare-text and `{`...`}` template-literal form is made
// from the snippet's *intent*, not from what's currently in the file:
//   - if the template has any interpolation, conditional, newline, or a
//     character JSX would interpret specially (`{`, `}`), wrap in `{`...`}`;
//   - otherwise emit bare text, with no JS escaping.
//
// Escaping for backticks/backslashes/${} only happens when the output is
// going to be inside a backtick literal. Bare JSX text doesn't interpret
// any of those, so escaping there would corrupt user-visible output.
func renderForJSXChild(tpl []render.Node, declared map[string]struct{}) string {
	if needsTemplateLiteral(tpl, declared) {
		return "{`" + render.RenderForLDApplicationTemplate(tpl, declared) + "`}"
	}
	bare, err := render.RenderForJSXText(tpl, declared)
	if err != nil {
		// Defensive: needsTemplateLiteral should have routed us to the
		// template path. If we somehow got here with interpolation, fall
		// back to the safe wrapping form.
		return "{`" + render.RenderForLDApplicationTemplate(tpl, declared) + "`}"
	}
	return bare
}

// declaredInputSet returns the set of input names declared on the snippet's
// frontmatter. Used to differentiate `{{ name }}` we own (declared inputs)
// from foreign template syntax (e.g. Vue's `{{ flagValue }}` mustaches in a
// Vue snippet body) that should pass through verbatim.
func declaredInputSet(s *model.Snippet) map[string]struct{} {
	out := make(map[string]struct{}, len(s.Frontmatter.Inputs))
	for name := range s.Frontmatter.Inputs {
		out[name] = struct{}{}
	}
	return out
}

// splitSurroundingWS returns the leading and trailing whitespace of s.
// If s is all whitespace, the leading captures it and trailing is empty.
func splitSurroundingWS(s string) (string, string) {
	leadEnd := 0
	for leadEnd < len(s) && isSpace(s[leadEnd]) {
		leadEnd++
	}
	if leadEnd == len(s) {
		return s, ""
	}
	trailStart := len(s)
	for trailStart > leadEnd && isSpace(s[trailStart-1]) {
		trailStart--
	}
	return s[:leadEnd], s[trailStart:]
}

func isSpace(b byte) bool { return b == ' ' || b == '\t' || b == '\n' || b == '\r' }

func needsTemplateLiteral(tpl []render.Node, declared map[string]struct{}) bool {
	if render.HasInterpolation(tpl, declared) {
		return true
	}
	for _, n := range tpl {
		switch x := n.(type) {
		case *render.Literal:
			if strings.Contains(x.Text, "\n") {
				return true
			}
			if render.ContainsJSXSpecial(x.Text) {
				return true
			}
		case *render.Var:
			// Foreign-template Vars (Vue/Handlebars `{{ name }}`) are
			// emitted verbatim. Their `{` and `}` would be interpreted
			// as JSX expression delimiters in bare-text mode, so force
			// the template-literal path so the curlies get escaped.
			return true
		}
	}
	return false
}
