package model

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// Frontmatter is the YAML block at the top of a .snippet.md file.
// ParseFile decodes with `KnownFields(true)`, so any frontmatter key not
// modelled below is rejected at parse time — that's deliberate, it catches
// typos like `Description:` or `entrypiont:` rather than silently dropping
// them. Adding a new schema field means adding it to this struct first.
type Frontmatter struct {
	ID          string           `yaml:"id"`
	SDK         string           `yaml:"sdk"`
	Kind        string           `yaml:"kind"`
	Lang        string           `yaml:"lang"`
	File        string           `yaml:"file"`
	Description string           `yaml:"description"`
	Inputs     map[string]Input `yaml:"inputs"`
	Validation Validation       `yaml:"validation"`
}

type Input struct {
	Type           string `yaml:"type"`
	Description    string `yaml:"description"`
	RuntimeDefault string `yaml:"runtime-default"`
}

type Validation struct {
	// Runtime selects the validator harness under
	// validators/languages/<runtime>/. If empty, the snippet's `lang:`
	// field is used as the fallback (e.g. lang=python implies the python
	// harness). Set explicitly when the snippet's lang doesn't equal the
	// runtime — e.g. `lang: javascript` snippets that run under Node use
	// `runtime: node`, and `lang: ts` snippets that build a browser bundle
	// for the JavaScript client SDK use `runtime: js-client`.
	Runtime string `yaml:"runtime"`

	// Entrypoint is the relative file path the harness invokes. If empty,
	// the snippet's `file:` field is used (which is also where the rendered
	// body is staged). Required for the validator to consider the snippet
	// runnable; absence means "no validation for this snippet."
	Entrypoint string `yaml:"entrypoint"`

	// Requirements is a runtime-specific dependency descriptor. For Python
	// it's the contents of requirements.txt. For other runtimes it's
	// language-specific (or empty when manifest companions carry deps).
	Requirements string `yaml:"requirements"`

	// Companions lists snippet IDs to stage alongside this one. Each
	// companion's body is rendered with the same runtime inputs and
	// written to the staging dir at the companion's `file:` path. Use for
	// multi-file projects (Java pom.xml, .NET .csproj, Rust Cargo.toml,
	// iOS Podfile, etc.). The companions need not declare their own
	// `validation.runtime` — they just declare `file:` so the validator
	// knows where to put them.
	Companions []string `yaml:"companions"`

	// Scaffold references another snippet (typically `kind: scaffold`)
	// whose body wraps this snippet's body via a `{{ body }}` placeholder.
	// Used by docs snippets that aren't standalone-runnable (a single
	// `variation()` call, a hook class definition, etc.) — the scaffold
	// supplies the surrounding init+context so a real validator run can
	// exercise the fragment. The scaffold owns the runtime, entrypoint,
	// requirements, and companions; this snippet contributes only the body
	// fragment plus any ScaffoldInputs.
	//
	// Mutually exclusive with Runtime/Entrypoint/Requirements/Companions
	// here — those come from the scaffold.
	Scaffold string `yaml:"scaffold"`

	// ScaffoldInputs supplies extra named values to the scaffold's
	// templating beyond the special `body` placeholder. Lets one scaffold
	// serve several wrappees that need slightly different setup (e.g.
	// pre-populated flag values, context attributes).
	ScaffoldInputs map[string]string `yaml:"scaffold-inputs"`
}

// Snippet pairs the frontmatter with the body of the first fenced code block
// in the markdown. The first-pass snippet file format is exactly "one fenced
// code block per file"; later phases may extend to multiple blocks per file.
type Snippet struct {
	Path        string
	Frontmatter Frontmatter
	CodeLang    string
	CodeBody    string
}

var frontmatterRe = regexp.MustCompile(`(?s)\A---\n(.*?)\n---\n`)
// `[ \t]*` (not `\s*`) for the trailing horizontal whitespace because
// `\s` includes `\n`: a greedy `\s*$` will consume the line-terminating
// newline, and then the `after = after[1:]` step in firstCodeBlock skips
// the *next* newline — silently dropping a blank line that immediately
// followed the fence. CommonMark says fenced code-block content
// preserves leading blank lines, so we have to too.
var fenceOpenRe = regexp.MustCompile("(?m)^```([a-zA-Z0-9_+-]*)[ \t]*$")
var fenceCloseRe = regexp.MustCompile("(?m)^```[ \t]*$")

// ParseFile reads a .snippet.md file from fsys and returns the parsed Snippet.
// snippetPath is an fs.FS-style path (slash-separated, no leading "./").
func ParseFile(fsys fs.FS, snippetPath string) (*Snippet, error) {
	raw, err := fs.ReadFile(fsys, snippetPath)
	if err != nil {
		return nil, err
	}
	m := frontmatterRe.FindSubmatchIndex(raw)
	if m == nil {
		return nil, fmt.Errorf("%s: missing YAML frontmatter", snippetPath)
	}
	var fm Frontmatter
	dec := yaml.NewDecoder(bytes.NewReader(raw[m[2]:m[3]]))
	dec.KnownFields(true)
	if err := dec.Decode(&fm); err != nil {
		return nil, fmt.Errorf("%s: frontmatter parse: %w", snippetPath, err)
	}
	body := raw[m[1]:]

	lang, code, err := firstCodeBlock(body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", snippetPath, err)
	}
	return &Snippet{
		Path:        snippetPath,
		Frontmatter: fm,
		CodeLang:    lang,
		CodeBody:    code,
	}, nil
}

// firstCodeBlock returns the language and body of the first fenced code block.
// The body does NOT include the trailing newline before the closing fence.
func firstCodeBlock(body []byte) (string, string, error) {
	openIdx := fenceOpenRe.FindSubmatchIndex(body)
	if openIdx == nil {
		return "", "", errors.New("no fenced code block found")
	}
	lang := string(body[openIdx[2]:openIdx[3]])
	// Content starts on the line after the opening fence.
	after := body[openIdx[1]:]
	if !bytes.HasPrefix(after, []byte("\n")) {
		return "", "", errors.New("malformed code fence: expected newline after opening fence")
	}
	after = after[1:]
	closeIdx := fenceCloseRe.FindIndex(after)
	if closeIdx == nil {
		return "", "", errors.New("unterminated fenced code block")
	}
	codeBytes := after[:closeIdx[0]]
	// Strip the final newline that precedes ```.
	codeBytes = bytes.TrimSuffix(codeBytes, []byte("\n"))
	return lang, string(codeBytes), nil
}

// LoadAll walks fsys for every *.snippet.md file and returns them indexed by id.
// fsys is rooted at the sdks/ directory's contents (so an entry path looks like
// "<sdk-id>/snippets/<group>/<name>.snippet.md").
func LoadAll(fsys fs.FS) (map[string]*Snippet, error) {
	out := map[string]*Snippet{}
	err := fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".snippet.md") {
			return nil
		}
		s, err := ParseFile(fsys, p)
		if err != nil {
			return err
		}
		if s.Frontmatter.ID == "" {
			return fmt.Errorf("%s: frontmatter.id is required", p)
		}
		if prev, ok := out[s.Frontmatter.ID]; ok {
			return fmt.Errorf("duplicate snippet id %q in %s and %s", s.Frontmatter.ID, prev.Path, p)
		}
		out[s.Frontmatter.ID] = s
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SortedIDs returns the ids in lexicographic order; useful for deterministic output.
func SortedIDs(snippets map[string]*Snippet) []string {
	ids := make([]string, 0, len(snippets))
	for id := range snippets {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
