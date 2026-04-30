package lddocs

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/markers"
	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
	"github.com/launchdarkly/sdk-meta/snippets/internal/version"
)

// markerSentinel is the substring every render marker contains. Used as a
// fast pre-filter so walking large doc trees only invokes the marker
// scanner on files that could possibly contain a marker.
var markerSentinel = []byte("SDK_SNIPPET:RENDER:")

// candidateExtensions is the set of doc file types the MDX scanner
// understands. ld-docs-private uses .mdx exclusively today; .md is
// included to cover ld-docs-public-style markdown that Fern compiles.
var candidateExtensions = map[string]struct{}{
	".mdx": {}, ".md": {},
}

// skipDirNames are directories we never descend into. Standard generated /
// dependency directories that never carry doc markers.
var skipDirNames = map[string]struct{}{
	"node_modules": {}, ".git": {}, ".next": {}, "dist": {}, "build": {},
	".cache": {}, "coverage": {}, "out": {}, ".turbo": {}, ".fern": {},
}

// Render walks every entrypoint, finds MDX files containing render markers,
// and rewrites each fenced code block whose marker references a snippet.
// Returns one entry per file it touched.
func Render(sdksFS fs.FS, entrypoints []string) ([]string, error) {
	return run(sdksFS, entrypoints, false)
}

// Verify re-renders every marked region in memory and fails if any hash in
// a marker does not match the bytes currently between its fences, or if a
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

// discoverFilesUnder walks every entrypoint, returning the set of doc files
// that contain the SDK_SNIPPET:RENDER: sentinel. Mirrors the
// ld-application adapter's discovery logic.
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
				if p != abs {
					if _, skip := skipDirNames[d.Name()]; skip {
						return filepath.SkipDir
					}
				}
				return nil
			}
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

// rewriteFile rewrites the fence bodies of every marked region. If dryRun
// is true it only verifies that (a) every marker carries a hash field,
// (b) the hash matches the current fence body, and (c) re-rendering would
// produce the same bytes.
func rewriteFile(path string, snippets map[string]*model.Snippet, dryRun bool) (bool, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return false, err
	}
	src := string(raw)

	matches, err := markers.ScanMDX(src)
	if err != nil {
		return false, err
	}
	if len(matches) == 0 {
		return false, nil
	}

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
		body, err := render.RenderRuntime(tpl, docsInputs(s))
		if err != nil {
			return false, fmt.Errorf("snippet %s: %w", s.Path, err)
		}

		// Hash covers ONLY the fence body bytes — the surrounding
		// <CodeBlocks>/<CodeBlock title='…'> wrappers and the fence
		// language tag are the consumer's choice.
		newHash := markers.HashContent(body)

		if dryRun {
			if m.Fields.Hash == "" {
				return false, fmt.Errorf("marker %q: missing required hash= field — re-render to populate it", m.Fields.ID)
			}
			actualHash := markers.HashContent(src[m.RegionStart:m.RegionEnd])
			if m.Fields.Hash != actualHash {
				return false, fmt.Errorf("marker %q: hand-edit detected — body hash %s does not match marker %s",
					m.Fields.ID, actualHash, m.Fields.Hash)
			}
			if body != src[m.RegionStart:m.RegionEnd] {
				return false, fmt.Errorf("marker %q: re-render would change region — run `snippets render`", m.Fields.ID)
			}
			continue
		}

		// Preserve the existing version stamp when the body is byte-
		// identical to what's on disk. Same semantics as the
		// ld-application adapter.
		ver := m.Fields.Version
		if ver == "" || body != src[m.RegionStart:m.RegionEnd] {
			ver = version.Version
		}
		newMarker := markers.FormatMDXMarker(markers.MarkerFields{
			ID:      m.Fields.ID,
			Hash:    newHash,
			Version: ver,
			Scope:   "content",
		})
		sb.WriteString(src[cursor:m.CommentStart])
		sb.WriteString(newMarker)
		sb.WriteString(src[m.CommentEnd:m.RegionStart])
		sb.WriteString(body)
		cursor = m.RegionEnd

		if src[m.CommentStart:m.CommentEnd] != newMarker || src[m.RegionStart:m.RegionEnd] != body {
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

// docsInputs synthesizes a value for each declared input. Bound inputs
// (those with a non-empty runtime-default) pass through; unbound
// credential-flavored inputs get a YOUR_<TYPE> placeholder so the doc
// renders human-friendly literals (`YOUR_SDK_KEY`, `YOUR_FLAG_KEY`)
// without leaking real keys. Plain string inputs use their runtime-default
// (which may be empty — empty values cause `{{ if name }}…{{ end }}`
// conditionals to omit, which is usually what the doc wants).
func docsInputs(s *model.Snippet) map[string]string {
	out := make(map[string]string, len(s.Frontmatter.Inputs))
	for name, in := range s.Frontmatter.Inputs {
		out[name] = docsDefault(name, in)
	}
	return out
}

// docsDefault picks the value substituted into the snippet for the docs
// render. The mapping is intentionally narrow: each input type that
// represents a credential gets a literal placeholder; plain strings fall
// through to the runtime default the snippet author specified.
func docsDefault(name string, in model.Input) string {
	switch in.Type {
	case "flag-key":
		return "YOUR_FLAG_KEY"
	case "sdk-key":
		return "YOUR_SDK_KEY"
	case "client-side-id":
		return "YOUR_CLIENT_SIDE_ID"
	case "mobile-key":
		return "YOUR_MOBILE_KEY"
	case "string":
		return in.RuntimeDefault
	default:
		// Unknown / future input type. Synthesize a screaming-snake
		// placeholder from the name so the doc reader sees a clear
		// "fill this in" rather than an unrendered template.
		return "YOUR_" + strings.ToUpper(toScreamingSnake(name))
	}
}

// toScreamingSnake converts camelCase / kebab-case / snake_case into
// SCREAMING_SNAKE_CASE for placeholder names.
func toScreamingSnake(s string) string {
	var sb strings.Builder
	for i, r := range s {
		switch {
		case r == '-' || r == '_':
			sb.WriteByte('_')
		case r >= 'A' && r <= 'Z':
			if i > 0 {
				sb.WriteByte('_')
			}
			sb.WriteRune(r)
		case r >= 'a' && r <= 'z':
			sb.WriteRune(r - 'a' + 'A')
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// atomicWriteFile is identical to ldapplication's atomic-rename path —
// duplicated rather than abstracted because the surface is small enough
// that a shared helper just adds an indirection.
func atomicWriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	mode := os.FileMode(0o644)
	if info, err := os.Stat(path); err == nil {
		mode = info.Mode().Perm()
	}
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
