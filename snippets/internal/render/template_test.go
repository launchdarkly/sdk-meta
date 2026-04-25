package render

import (
	"strings"
	"testing"
)

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
	got = RenderForLDApplicationTemplate(nodes)
	want = "echo \"launchdarkly-server-sdk${version ? `==${version}` : ''}\" done"
	if got != want {
		t.Fatalf("ld-application: got %q want %q", got, want)
	}
}

func TestRenderForLDApplicationTemplateEscapes(t *testing.T) {
	nodes, _ := Parse("a \\ b ` c ${d} {{ name }}")
	got := RenderForLDApplicationTemplate(nodes)
	want := "a \\\\ b \\` c \\${d} ${name}"
	if got != want {
		t.Fatalf("escapes: got %q want %q", got, want)
	}
}

// Regression for review #4: a variable starting with `end` previously hit a
// HasPrefix check and was treated as a block-close. Now the case requires
// equality so `endTime` is recognized as a normal variable.
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

func TestRenderRuntimeUnknownVar(t *testing.T) {
	nodes, _ := Parse(`{{ missing }}`)
	_, err := RenderRuntime(nodes, map[string]string{})
	if err == nil || !strings.Contains(err.Error(), "missing runtime input") {
		t.Fatalf("want missing-input error, got %v", err)
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

// Regression for review #10: the bare-JSX-text path must not apply
// template-literal escaping. A backslash in the source should round-trip
// verbatim.
func TestRenderForJSXTextNoEscape(t *testing.T) {
	nodes, _ := Parse(`python .\main.py`)
	got, err := RenderForJSXText(nodes)
	if err != nil {
		t.Fatalf("render: %v", err)
	}
	if got != `python .\main.py` {
		t.Fatalf("backslash mangled: %q", got)
	}
}

func TestRenderForJSXTextRefusesInterp(t *testing.T) {
	nodes, _ := Parse(`hello {{ name }}`)
	if _, err := RenderForJSXText(nodes); err == nil {
		t.Fatalf("want error when template has interpolation")
	}
}

func TestHasInterpolation(t *testing.T) {
	cases := map[string]bool{
		"plain":              false,
		"hello {{ name }}":   true,
		"{{ if a }}b{{ end }}": true,
	}
	for src, want := range cases {
		nodes, _ := Parse(src)
		if got := HasInterpolation(nodes); got != want {
			t.Errorf("HasInterpolation(%q) = %v, want %v", src, got, want)
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
