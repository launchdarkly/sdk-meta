package render

import "testing"

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
	got = RenderForLDApplication(nodes)
	want = "echo \"launchdarkly-server-sdk${version ? `==${version}` : ''}\" done"
	if got != want {
		t.Fatalf("ld-application: got %q want %q", got, want)
	}
}

func TestRenderForLDApplicationEscapes(t *testing.T) {
	nodes, _ := Parse("a \\ b ` c ${d} {{ name }}")
	got := RenderForLDApplication(nodes)
	want := "a \\\\ b \\` c \\${d} ${name}"
	if got != want {
		t.Fatalf("escapes: got %q want %q", got, want)
	}
}
