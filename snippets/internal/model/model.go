package model

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// Frontmatter is the YAML block at the top of a .snippet.md file.
// Only the fields this first-pass slice actually uses are modelled; unknown
// fields are ignored so later phases can extend the schema without churn here.
type Frontmatter struct {
	ID          string           `yaml:"id"`
	SDK         string           `yaml:"sdk"`
	Kind        string           `yaml:"kind"`
	Lang        string           `yaml:"lang"`
	File        string           `yaml:"file"`
	Description string           `yaml:"description"`
	Inputs        map[string]Input  `yaml:"inputs"`
	LDApplication LDApplicationHints `yaml:"ld-application"`
	Validation    Validation        `yaml:"validation"`
}

type Input struct {
	Type           string `yaml:"type"`
	Description    string `yaml:"description"`
	RuntimeDefault string `yaml:"runtime-default"`
}

type LDApplicationHints struct {
	Slot string `yaml:"slot"`
}

type Validation struct {
	Entrypoint   string `yaml:"entrypoint"`
	Requirements string `yaml:"requirements"`
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
var fenceOpenRe = regexp.MustCompile("(?m)^```([a-zA-Z0-9_+-]*)\\s*$")
var fenceCloseRe = regexp.MustCompile("(?m)^```\\s*$")

// ParseFile reads a .snippet.md file and returns the parsed Snippet.
func ParseFile(path string) (*Snippet, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	m := frontmatterRe.FindSubmatchIndex(raw)
	if m == nil {
		return nil, fmt.Errorf("%s: missing YAML frontmatter", path)
	}
	var fm Frontmatter
	if err := yaml.Unmarshal(raw[m[2]:m[3]], &fm); err != nil {
		return nil, fmt.Errorf("%s: frontmatter parse: %w", path, err)
	}
	body := raw[m[1]:]

	lang, code, err := firstCodeBlock(body)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &Snippet{
		Path:        path,
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

// LoadAll walks sdksDir for every *.snippet.md file and returns them indexed by id.
func LoadAll(sdksDir string) (map[string]*Snippet, error) {
	out := map[string]*Snippet{}
	err := filepath.WalkDir(sdksDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".snippet.md") {
			return nil
		}
		s, err := ParseFile(p)
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
