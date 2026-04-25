package ldapplication

import (
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
func discoverTargetFiles(sdksDir, appDir string) ([]string, error) {
	entries, err := os.ReadDir(sdksDir)
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
		if desc.LDApplication.GetStartedFile == "" {
			continue
		}
		full := filepath.Join(appDir, desc.LDApplication.GetStartedFile)
		if _, err := os.Stat(full); err != nil {
			return nil, fmt.Errorf("descriptor %s: target not found: %w", descPath, err)
		}
		out = append(out, full)
	}
	return out, nil
}

// rewriteFile does the actual in-place rewrite. If dryRun is true it only
// verifies that (a) every marker hash matches the current file content and
// (b) re-rendering produces the same bytes.
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
		rendered := render.RenderForLDApplication(tpl)
		jsxBody := wrapForJSX(rendered, src, m.RegionStart, m.RegionEnd)

		if dryRun {
			actualHash := markers.HashContent(src[m.RegionStart:m.RegionEnd])
			if m.Fields.Hash != "" && m.Fields.Hash != actualHash {
				return false, fmt.Errorf("marker %q: hand-edit detected — file hash %s does not match marker %s",
					m.Fields.ID, actualHash, m.Fields.Hash)
			}
			if jsxBody != src[m.RegionStart:m.RegionEnd] {
				return false, fmt.Errorf("marker %q: re-render would change region — run `snippets render`", m.Fields.ID)
			}
			continue
		}

		newMarker := markers.FormatMarker(m.Style, markers.MarkerFields{
			ID:      m.Fields.ID,
			Hash:    markers.HashContent(jsxBody),
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
	return true, os.WriteFile(path, []byte(sb.String()), 0o644)
}

// wrapForJSX produces the bytes that belong between <Tag> and </Tag>.
// To keep cut-over diffs minimal, it preserves the surrounding-whitespace
// shape of the existing region:
//   - leading whitespace (between `>` and the first non-space char) is reused
//   - trailing whitespace (between the last non-space char and `</`) is reused
//
// The middle is replaced. Bare-text vs. backtick-template form is decided by
// whether the rendered content needs interpolation/multiline.
func wrapForJSX(rendered, src string, regionStart, regionEnd int) string {
	hasInterp := strings.Contains(rendered, "${")
	isMultiline := strings.Contains(rendered, "\n")
	needsTemplate := hasInterp || isMultiline

	existing := src[regionStart:regionEnd]
	leading, trailing := splitSurroundingWS(existing)

	wasBare := !strings.Contains(strings.TrimSpace(existing), "{`") &&
		!strings.HasPrefix(strings.TrimSpace(existing), "{")

	if wasBare && !needsTemplate {
		return leading + rendered + trailing
	}
	return leading + "{`" + rendered + "`}" + trailing
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
