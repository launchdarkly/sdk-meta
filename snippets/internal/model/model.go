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
	Inputs      map[string]Input `yaml:"inputs"`
	Validation  Validation       `yaml:"validation"`
}

type Input struct {
	Type           string `yaml:"type"`
	Description    string `yaml:"description"`
	RuntimeDefault string `yaml:"runtime-default"`
}

type Validation struct {
	// Checks is the list of checks to run against this snippet. Each check
	// is a Validator invocation with its own kind, scaffold, runtime, and
	// env. The dispatcher runs them sequentially; a single failure stops
	// the snippet (failed checks short-circuit later ones).
	//
	// When Checks is non-empty it is authoritative — the top-level
	// `runtime/scaffold/entrypoint/...` fields are taken as defaults that
	// each check inherits unless it overrides. When Checks is empty BUT
	// any of the top-level legacy fields is set, the loader synthesizes
	// a single legacy Check `[{Kind: "runtime", ...top-level fields}]`.
	// This keeps every existing snippet's behavior identical without
	// requiring schema migration.
	//
	// Recognized Check.Kind values:
	//   - "parse"     — fast syntactic check, ideally just the language's
	//                   built-in parser (node --check, php -l, ruby -wc,
	//                   etc.). Doesn't need LD env. Catches reserved-word
	//                   parameter names, missing semicolons, broken
	//                   fences. Default for `kind: reference` snippets
	//                   that declare a scaffold but no explicit checks.
	//   - "typecheck" — language-level type / compile check that catches
	//                   undefined identifiers and type errors but does
	//                   not run the program (tsc --noEmit, dart analyze,
	//                   xcodebuild build, ./gradlew compileDebugKotlin,
	//                   clang -fsyntax-only). Per-language opt-in:
	//                   harnesses that don't ship a typechecker treat
	//                   the kind as `parse`. Doesn't need LD env.
	//   - "runtime"   — full end-to-end execution against a real
	//                   LaunchDarkly environment, asserting on the
	//                   EXAM-HELLO success line. Existing single-check
	//                   snippets are all `runtime` for back-compat.
	//
	// The dispatcher forwards `SNIPPET_CHECK=<kind>` to the harness so
	// `validators/languages/<runtime>/harness/run.sh` can switch on it.
	// A harness that doesn't recognize the kind should exit non-zero
	// with a clear "unsupported check" message rather than silently
	// running its default path.
	Checks []Check `yaml:"checks,omitempty"`

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

	// Env maps validator-specific env-var names to literal values that
	// the dispatcher forwards into the harness's process environment.
	// Used by per-snippet binding metadata that the harness reads at
	// run time — e.g. ios-install's `INSTALL_KIND` discriminator that
	// picks `swift-package` vs `podfile` vs `cartfile`.
	//
	// Keys are restricted to a snippet-author-defined set; the
	// dispatcher does not touch the value (no env lookup, no template
	// substitution). Use Placeholders for the env-derived case.
	Env map[string]string `yaml:"env"`

	// Placeholders maps literal source-text fragments inside the snippet
	// body to environment-variable names. After the body is rendered (and
	// after any scaffold composition), the dispatcher does a literal
	// string-replace: every occurrence of the key is replaced with the
	// value of the named env var.
	//
	// Use case: gonfalon's sdk-info init snippets carry literal
	// `'YOUR_SDK_KEY'` / `'YOUR_MOBILE_KEY'` / `'YOUR_CLIENT_SIDE_ID'`
	// placeholders that the user is expected to swap manually. Validating
	// the snippet end-to-end against a real LaunchDarkly env requires
	// substituting a real key — but rewriting the body to use a
	// `{{ sdk_key }}` template marker would make the rendered output
	// (consumed by gonfalon) lose the human-readable placeholder.
	// Placeholders keep the snippet body unchanged at render time and
	// only swap at validate time.
	//
	// Only a small allow-list of env var names is honored:
	// LAUNCHDARKLY_SDK_KEY, LAUNCHDARKLY_FLAG_KEY, LAUNCHDARKLY_MOBILE_KEY,
	// LAUNCHDARKLY_CLIENT_SIDE_ID. The dispatcher rejects any other name.
	Placeholders map[string]string `yaml:"placeholders"`
}

// Check describes one validator invocation. A snippet's `validation.checks:`
// is a list of these; the dispatcher runs them in order. Every field except
// Kind is optional — unset fields inherit from the parent Validation, which
// is how a snippet with multiple checks sharing a scaffold avoids
// duplicating the scaffold name across entries.
type Check struct {
	// Kind picks the harness dispatch branch. One of "parse", "typecheck",
	// "runtime". The dispatcher forwards this verbatim as the
	// `SNIPPET_CHECK` env var. See the `Validation.Checks` doc comment
	// for the full enum description.
	Kind string `yaml:"kind"`

	// Scaffold, Runtime, Entrypoint, Companions, Requirements,
	// ScaffoldInputs, Env, and Placeholders mirror the parent Validation
	// fields and, when set on a Check, override the parent for this
	// check only. Common case is one shared scaffold for parse + runtime
	// checks; declare it on the parent Validation and leave it unset on
	// each Check.
	Scaffold       string            `yaml:"scaffold,omitempty"`
	Runtime        string            `yaml:"runtime,omitempty"`
	Entrypoint     string            `yaml:"entrypoint,omitempty"`
	Companions     []string          `yaml:"companions,omitempty"`
	Requirements   string            `yaml:"requirements,omitempty"`
	ScaffoldInputs map[string]string `yaml:"scaffold-inputs,omitempty"`
	Env            map[string]string `yaml:"env,omitempty"`
	Placeholders   map[string]string `yaml:"placeholders,omitempty"`
}

// EffectiveChecks returns the list of checks to run against this snippet's
// validation block, with each check's fields fully resolved (Check overrides
// followed by Validation defaults). When the snippet declares no explicit
// Checks but has legacy top-level fields, it synthesizes a single
// `kind: runtime` check from those fields — preserving every existing
// snippet's behavior under the multi-check dispatcher.
//
// Returns an empty slice when neither Checks nor any legacy field is set —
// the dispatcher treats that as "not validatable" (same as the legacy
// `isValidatable` predicate).
func (v Validation) EffectiveChecks() []Check {
	merge := func(c Check) Check {
		if c.Scaffold == "" {
			c.Scaffold = v.Scaffold
		}
		if c.Runtime == "" {
			c.Runtime = v.Runtime
		}
		if c.Entrypoint == "" {
			c.Entrypoint = v.Entrypoint
		}
		if len(c.Companions) == 0 {
			c.Companions = v.Companions
		}
		if c.Requirements == "" {
			c.Requirements = v.Requirements
		}
		if len(c.ScaffoldInputs) == 0 {
			c.ScaffoldInputs = v.ScaffoldInputs
		}
		// Env and Placeholders merge (check overrides win key-by-key) so a
		// check can add or override one env var without losing the parent's
		// others.
		if len(v.Env) > 0 {
			merged := map[string]string{}
			for k, val := range v.Env {
				merged[k] = val
			}
			for k, val := range c.Env {
				merged[k] = val
			}
			c.Env = merged
		}
		if len(v.Placeholders) > 0 {
			merged := map[string]string{}
			for k, val := range v.Placeholders {
				merged[k] = val
			}
			for k, val := range c.Placeholders {
				merged[k] = val
			}
			c.Placeholders = merged
		}
		return c
	}

	if len(v.Checks) > 0 {
		out := make([]Check, 0, len(v.Checks))
		for _, c := range v.Checks {
			out = append(out, merge(c))
		}
		return out
	}
	// Legacy fall-through: synthesize a single runtime check from the
	// top-level fields if any of them is set.
	if v.Runtime != "" || v.Entrypoint != "" || v.Scaffold != "" {
		return []Check{merge(Check{Kind: "runtime"})}
	}
	return nil
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
