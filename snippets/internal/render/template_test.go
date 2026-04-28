package render

import (
	"strings"
	"testing"
)

// allInputs returns a declaredInputs set containing the given names; helper
// for tests that don't care about foreign-template pass-through.
func allInputs(names ...string) map[string]struct{} {
	m := make(map[string]struct{}, len(names))
	for _, n := range names {
		m[n] = struct{}{}
	}
	return m
}

func TestParseAndRender(t *testing.T) {
	nodes, err := Parse(`echo "launchdarkly-server-sdk{{ if version }}=={{ version }}{{ end }}" done`)
	if err != nil {
		t.Fatal(err)
	}

	// Runtime, version set
	got, err := RenderRuntime(nodes, map[string]string{"version": "9.2.0"})
	if err != nil {
		t.Fatal(err)
	}
	want := `echo "launchdarkly-server-sdk==9.2.0" done`
	if got != want {
		t.Fatalf("runtime: got %q want %q", got, want)
	}

	// Runtime, version empty → conditional omitted
	got, err = RenderRuntime(nodes, map[string]string{"version": ""})
	if err != nil {
		t.Fatal(err)
	}
	want = `echo "launchdarkly-server-sdk" done`
	if got != want {
		t.Fatalf("runtime empty: got %q want %q", got, want)
	}

	// ld-application rendering produces a JS ternary expression
	got = RenderForLDApplicationTemplate(nodes, allInputs("version"))
	want = "echo \"launchdarkly-server-sdk${version ? `==${version}` : ''}\" done"
	if got != want {
		t.Fatalf("ld-application: got %q want %q", got, want)
	}
}

func TestRenderForLDApplicationTemplateEscapes(t *testing.T) {
	nodes, _ := Parse("a \\ b ` c ${d} {{ name }}")
	got := RenderForLDApplicationTemplate(nodes, allInputs("name"))
	want := "a \\\\ b \\` c \\${d} ${name}"
	if got != want {
		t.Fatalf("escapes: got %q want %q", got, want)
	}
}

// Regression: a variable starting with `end` previously hit a HasPrefix check
// and was treated as a block-close. Now equality is required.
func TestEndPrefixedVariable(t *testing.T) {
	nodes, err := Parse(`{{ endTime }}`)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	got, err := RenderRuntime(nodes, map[string]string{"endTime": "23:59"})
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if got != "23:59" {
		t.Fatalf("got %q want 23:59", got)
	}
}

func TestUnmatchedEnd(t *testing.T) {
	_, err := Parse(`hello {{ end }}`)
	if err == nil || !strings.Contains(err.Error(), "unmatched") {
		t.Fatalf("want unmatched-end error, got %v", err)
	}
}

func TestUnclosedIf(t *testing.T) {
	_, err := Parse(`{{ if v }}body without end`)
	if err == nil || !strings.Contains(err.Error(), "unclosed") {
		t.Fatalf("want unclosed-if error, got %v", err)
	}
}

// Foreign-template pass-through: an undeclared name in `{{ name }}` is
// emitted verbatim by RenderRuntime so a Vue snippet's `{{ flagValue }}`
// survives validation untouched.
func TestRenderRuntimePassesThroughUnknownVar(t *testing.T) {
	nodes, _ := Parse(`hello {{ flagValue }}`)
	got, err := RenderRuntime(nodes, map[string]string{})
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if got != `hello {{ flagValue }}` {
		t.Fatalf("got %q", got)
	}
}

// A conditional referring to an undeclared input is still an authoring bug,
// not a foreign template — Vue uses `v-if`, not `{{ if … }}`.
func TestRenderRuntimeRejectsUnknownCondVar(t *testing.T) {
	nodes, _ := Parse(`{{ if missing }}body{{ end }}`)
	_, err := RenderRuntime(nodes, map[string]string{})
	if err == nil || !strings.Contains(err.Error(), "undeclared input") {
		t.Fatalf("want undeclared-conditional error, got %v", err)
	}
}

func TestRenderRuntimeEmptyTemplate(t *testing.T) {
	nodes, err := Parse("")
	if err != nil {
		t.Fatal(err)
	}
	got, err := RenderRuntime(nodes, nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != "" {
		t.Fatalf("want empty, got %q", got)
	}
}

// Bare-JSX-text path must not apply template-literal escaping. A backslash
// in the source should round-trip verbatim.
func TestRenderForJSXTextNoEscape(t *testing.T) {
	nodes, _ := Parse(`python .\main.py`)
	got, err := RenderForJSXText(nodes, allInputs())
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if got != `python .\main.py` {
		t.Fatalf("backslash mangled: %q", got)
	}
}

// Declared interpolation in JSX-text rendering should fail loudly so the
// caller routes to the template-literal path.
func TestRenderForJSXTextRefusesDeclaredInterp(t *testing.T) {
	nodes, _ := Parse(`hello {{ name }}`)
	if _, err := RenderForJSXText(nodes, allInputs("name")); err == nil {
		t.Fatalf("want error when template has declared interpolation")
	}
}

// Foreign-template `{{ name }}` in JSX text is fine — passes through.
func TestRenderForJSXTextPassesForeignTemplate(t *testing.T) {
	nodes, _ := Parse(`Feature flag {{ flagValue }} reads as expected`)
	got, err := RenderForJSXText(nodes, allInputs())
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if got != `Feature flag {{ flagValue }} reads as expected` {
		t.Fatalf("got %q", got)
	}
}

func TestHasInterpolation(t *testing.T) {
	cases := []struct {
		src      string
		declared []string
		want     bool
	}{
		{"plain", nil, false},
		{"hello {{ name }}", []string{"name"}, true},
		{"hello {{ name }}", nil, false}, // foreign template, not interpolation
		{"{{ if a }}b{{ end }}", []string{"a"}, true},
	}
	for _, c := range cases {
		nodes, _ := Parse(c.src)
		if got := HasInterpolation(nodes, allInputs(c.declared...)); got != c.want {
			t.Errorf("HasInterpolation(%q, declared=%v) = %v, want %v", c.src, c.declared, got, c.want)
		}
	}
}

func TestContainsJSXSpecial(t *testing.T) {
	if !ContainsJSXSpecial("a { b") {
		t.Fatal("expected true for `{`")
	}
	if !ContainsJSXSpecial("a } b") {
		t.Fatal("expected true for `}`")
	}
	if ContainsJSXSpecial("plain text") {
		t.Fatal("expected false for plain text")
	}
}
