// Package rawfiles renders snippet bodies to plain text files outside any
// JSX/MDX context. It exists to back the gonfalon `packages/sdk-info/`
// surface, where the existing consumer pattern is `import x from
// './foo.txt?raw'` rather than the marker-driven `<Snippet>{...}</Snippet>`
// rewrite the ldapplication adapter performs.
//
// Unlike ldapplication and lddocs, this adapter does NOT scan consumer
// files for SDK_SNIPPET:RENDER markers. Marker-based discovery doesn't
// fit the `?raw` import pattern (the consumer references a path, not an
// inline marker), so the adapter is driven by a manifest the consumer
// commits next to the output directory. The manifest enumerates every
// snippet ID it wants extracted and the relative output path it should
// land at.
//
// There is no `verify` counterpart for this target: with no marker on the
// consumer side there's no contract to verify against. CI in the consumer
// repo runs the renderer and diffs the resulting tree.
package rawfiles

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/launchdarkly/sdk-meta/snippets/internal/model"
	"github.com/launchdarkly/sdk-meta/snippets/internal/render"
	"gopkg.in/yaml.v3"
)

// Manifest describes which snippets to extract and where to write them.
// It's authored by the consumer and committed alongside the output
// directory so the mapping from canonical snippet IDs to consumer paths
// is reviewable and version-controlled.
type Manifest struct {
	// Out is the output root, relative to the manifest's directory or
	// (less commonly) absolute. Each Files entry's Path is taken
	// relative to Out.
	Out string `yaml:"out"`

	// Files lists the (snippet-id, output-path) pairs to render. Order
	// in the source manifest is preserved; duplicate IDs and duplicate
	// output paths are both errors caught by Validate.
	Files []ManifestEntry `yaml:"files"`
}

// ManifestEntry is one line in the manifest. Both fields are required.
type ManifestEntry struct {
	// ID matches a snippet's frontmatter `id:` exactly. Mismatches are
	// rejected with a list of close-match suggestions so a typo doesn't
	// silently get skipped.
	ID string `yaml:"id"`

	// Path is the file location relative to Manifest.Out. The renderer
	// creates parent directories as needed.
	Path string `yaml:"path"`
}

// LoadManifest reads and parses a manifest file. Returns the parsed
// Manifest and the directory containing the manifest (used to resolve a
// relative `out:` path against).
func LoadManifest(manifestPath string) (*Manifest, string, error) {
	abs, err := filepath.Abs(manifestPath)
	if err != nil {
		return nil, "", fmt.Errorf("manifest %q: %w", manifestPath, err)
	}
	raw, err := os.ReadFile(abs)
	if err != nil {
		return nil, "", fmt.Errorf("manifest %q: %w", manifestPath, err)
	}
	var m Manifest
	dec := yaml.NewDecoder(strings.NewReader(string(raw)))
	dec.KnownFields(true)
	if err := dec.Decode(&m); err != nil {
		return nil, "", fmt.Errorf("manifest %q: %w", manifestPath, err)
	}
	if err := m.validate(); err != nil {
		return nil, "", fmt.Errorf("manifest %q: %w", manifestPath, err)
	}
	return &m, filepath.Dir(abs), nil
}

// validate enforces structural invariants on a parsed manifest:
//   - `out:` is required (otherwise the renderer wouldn't know where to write);
//   - every entry has both `id:` and `path:`;
//   - no duplicate ids (the consumer would silently shadow one);
//   - no duplicate output paths (two snippets racing to the same file is
//     either a copy-paste mistake or deliberate aliasing — neither is OK
//     in a generated tree).
func (m *Manifest) validate() error {
	if strings.TrimSpace(m.Out) == "" {
		return fmt.Errorf("`out:` is required")
	}
	if len(m.Files) == 0 {
		return fmt.Errorf("`files:` is empty")
	}
	seenID := map[string]struct{}{}
	seenPath := map[string]struct{}{}
	for i, e := range m.Files {
		if strings.TrimSpace(e.ID) == "" {
			return fmt.Errorf("files[%d]: `id:` is required", i)
		}
		if strings.TrimSpace(e.Path) == "" {
			return fmt.Errorf("files[%d] (id=%q): `path:` is required", i, e.ID)
		}
		// Output paths are always relative to `out:`. An absolute or
		// `..`-traversing path would let a manifest write outside the
		// declared output root, which defeats the point of having one.
		if filepath.IsAbs(e.Path) {
			return fmt.Errorf("files[%d] (id=%q): `path:` must be relative, got %q", i, e.ID, e.Path)
		}
		clean := filepath.Clean(e.Path)
		if strings.HasPrefix(clean, "..") || strings.Contains(clean, string(filepath.Separator)+"..") {
			return fmt.Errorf("files[%d] (id=%q): `path:` must not escape `out:`, got %q", i, e.ID, e.Path)
		}
		if _, dup := seenID[e.ID]; dup {
			return fmt.Errorf("duplicate id %q in manifest", e.ID)
		}
		seenID[e.ID] = struct{}{}
		if _, dup := seenPath[clean]; dup {
			return fmt.Errorf("duplicate output path %q in manifest", e.Path)
		}
		seenPath[clean] = struct{}{}
	}
	return nil
}

// Render extracts every snippet listed in the manifest at manifestPath
// and writes its rendered body to <consumerDir>/<manifest.out>/<entry.path>.
// `consumerDir` is the consumer-checkout root that `out:` is resolved
// against; passing the empty string defaults to the manifest's own
// directory (the common case where the manifest sits next to the output
// directory).
//
// Returns the absolute paths of every file written, sorted lexically.
// Files whose content is byte-identical to what's already on disk are
// still listed — atomic writes mean the inode changes regardless, and
// the caller (CI's `git diff`) is the right place to detect no-op runs.
func Render(sdksFS fs.FS, manifestPath, consumerDir string) ([]string, error) {
	m, manifestDir, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, err
	}

	snippets, err := model.LoadAll(sdksFS)
	if err != nil {
		return nil, err
	}

	root := consumerDir
	if root == "" {
		root = manifestDir
	}
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("consumer dir %q: %w", root, err)
	}
	outRoot := m.Out
	if !filepath.IsAbs(outRoot) {
		outRoot = filepath.Join(absRoot, outRoot)
	}

	var written []string
	for _, e := range m.Files {
		s, ok := snippets[e.ID]
		if !ok {
			return nil, fmt.Errorf("manifest references unknown snippet id %q (suggestions: %s)",
				e.ID, suggestSimilarIDs(e.ID, snippets))
		}
		body, err := renderBody(s)
		if err != nil {
			return nil, fmt.Errorf("snippet %s: %w", s.Path, err)
		}
		dest := filepath.Join(outRoot, filepath.FromSlash(e.Path))
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return nil, fmt.Errorf("create %s: %w", filepath.Dir(dest), err)
		}
		if err := atomicWriteFile(dest, []byte(body)); err != nil {
			return nil, fmt.Errorf("write %s: %w", dest, err)
		}
		written = append(written, dest)
	}
	return written, nil
}

// renderBody runs the snippet's body through the runtime renderer with
// no inputs. Most sdk-info bodies have no template inputs at all, so
// this is a near-passthrough; declared inputs would render as empty
// strings here, and undeclared `{{ name }}` syntax (e.g. the cursor
// prompt's `{{SDK_NAME}}` runtime placeholders) round-trips as literal
// `{{ name }}` per the renderer's foreign-template contract.
//
// The output always ends with a trailing newline so the written file
// matches the POSIX-friendly convention the source `.txt` files used.
// The .snippet.md fence syntax strips the trailing newline before the
// closing fence; we add one back here.
func renderBody(s *model.Snippet) (string, error) {
	tpl, err := render.Parse(s.CodeBody)
	if err != nil {
		return "", err
	}
	body, err := render.RenderRuntime(tpl, nil)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(body, "\n") {
		body += "\n"
	}
	return body, nil
}

// suggestSimilarIDs returns a short comma-joined list of candidate IDs
// that look similar to the missing one. Used in the manifest-mismatch
// error so a typo is easy to fix without paging through a 300-line
// list of every loaded snippet.
func suggestSimilarIDs(missing string, snippets map[string]*model.Snippet) string {
	const max = 5
	var hits []string
	parts := strings.Split(missing, "/")
	tail := parts[len(parts)-1]
	for id := range snippets {
		if strings.Contains(id, tail) {
			hits = append(hits, id)
			if len(hits) >= max {
				break
			}
		}
	}
	if len(hits) == 0 {
		return "(none)"
	}
	return strings.Join(hits, ", ")
}

// atomicWriteFile mirrors ldapplication.atomicWriteFile: write to a
// same-directory tempfile, fsync, rename. Crash-safe and concurrent-safe
// across parallel renders that target distinct destinations.
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
