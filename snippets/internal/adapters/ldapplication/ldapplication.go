package ldapplication

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/markers"
	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
	"github.com/launchdarkly/sdk-meta/snippets/internal/version"
)

// Render walks every SDK's get-started TSX file under appDir, finds render
// markers, and rewrites each marked region with the rendered snippet content.
// Returns one entry per file it touched.
func Render(sdksDir, appDir string) ([]string, error) {
	snippets, err := model.LoadAll(sdksDir)
	if err != nil {
		return nil, err
	}

	files, err := discoverTargetFiles(sdksDir, appDir)
	if err != nil {
		return nil, err
	}

	var changed []string
	for _, path := range files {
		ok, err := rewriteFile(path, snippets, false)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if ok {
			changed = append(changed, path)
		}
	}
	return changed, nil
}

// Verify re-renders every marked region in memory and fails if any hash in a
// marker does not match the hash of the bytes currently in the file, or if a
// re-render would change content. Never modifies files.
func Verify(sdksDir, appDir string) error {
	snippets, err := model.LoadAll(sdksDir)
	if err != nil {
		return err
	}

	files, err := discoverTargetFiles(sdksDir, appDir)
	if err != nil {
		return err
	}

	for _, path := range files {
		if _, err := rewriteFile(path, snippets, true); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
	}
	return nil
}

// discoverTargetFiles returns every file referenced by an sdk.yaml's
// `ld-application.get-started-file` field, resolved relative to appDir.
//
// Each referenced path must be a clean relative path that stays inside
// appDir. This guards against a malicious sdk.yaml committing
// `get-started-file: ../../../foo` and the renderer overwriting arbitrary
// files outside the consumer checkout.
func discoverTargetFiles(sdksDir, appDir string) ([]string, error) {
	entries, err := os.ReadDir(sdksDir)
	if err != nil {
		return nil, err
	}
	absAppDir, err := filepath.Abs(appDir)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		descPath := filepath.Join(sdksDir, e.Name(), "sdk.yaml")
		desc, err := loadDescriptor(descPath)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		rel := desc.LDApplication.GetStartedFile
		if rel == "" {
			continue
		}
		if filepath.IsAbs(rel) {
			return nil, fmt.Errorf("descriptor %s: get-started-file %q must be relative", descPath, rel)
		}
		full := filepath.Join(absAppDir, rel)
		// Reject any path that escapes appDir. filepath.Rel followed by a
		// `..` prefix check is the canonical way to do this.
		relCheck, err := filepath.Rel(absAppDir, full)
		if err != nil || relCheck == ".." || strings.HasPrefix(relCheck, ".."+string(filepath.Separator)) {
			return nil, fmt.Errorf("descriptor %s: get-started-file %q escapes appDir", descPath, rel)
		}
		if _, err := os.Stat(full); err != nil {
			return nil, fmt.Errorf("descriptor %s: target not found: %w", descPath, err)
		}
		out = append(out, full)
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

		newMarker := markers.FormatMarker(m.Style, markers.MarkerFields{
			ID:      m.Fields.ID,
			Hash:    newHash,
			Version: version.Version,
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
	return true, atomicWriteFile(path, []byte(sb.String()))
}

// atomicWriteFile writes to a same-directory tempfile, fsyncs, and renames
// over the destination. The destination's permission bits are preserved so
// running `snippets render` on a checkout that has tightened permissions
// (e.g. read-only mode for a CODEOWNER-protected file) doesn't quietly
// reset them to 0644.
func atomicWriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	mode := os.FileMode(0o644)
	if info, err := os.Stat(path); err == nil {
		mode = info.Mode().Perm()
	}
	// Random suffix avoids colliding with parallel renders.
	var sfx [8]byte
	if _, err := rand.Read(sfx[:]); err != nil {
		return err
	}
	tmp := filepath.Join(dir, "."+filepath.Base(path)+".sdk-snippets."+hex.EncodeToString(sfx[:])+".tmp")
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Sync(); err != nil {
		f.Close()
		os.Remove(tmp)
		return err
	}
	if err := f.Close(); err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
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
