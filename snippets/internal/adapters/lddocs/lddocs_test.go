package lddocs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeFile is a small fixture helper for parallel tests.
func writeFile(t *testing.T, dir, rel, body string) {
	t.Helper()
	p := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

const installSnippet = `---
id: test-sdk/install
sdk: test-sdk
kind: install
lang: shell
description: Install with optional version pin.
inputs:
  version:
    type: string
    description: Optional pinned version
    runtime-default: ""
---

Install:

` + "```shell\n" +
	`pip install foo{{ if version }}=={{ version }}{{ end }}` + "\n" +
	`` + "```\n"

const initSnippet = `---
id: test-sdk/init
sdk: test-sdk
kind: hello-world
lang: python
file: main.py
description: Initialize.
inputs:
  flagKey:
    type: flag-key
    description: Flag to evaluate.
---

Init:

` + "```python\n" +
	`import foo` + "\n" +
	`flag = foo.variation("{{ flagKey }}")` + "\n" +
	`` + "```\n"

// Render fills hash=0 placeholders with real content + hash, and Verify
// then accepts the resulting file.
func TestRenderThenVerify_RoundTrip(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	writeFile(t, sdks, "test-sdk/sdk.yaml", "id: test-sdk\n")
	writeFile(t, sdks, "test-sdk/snippets/install.snippet.md", installSnippet)
	writeFile(t, sdks, "test-sdk/snippets/init.snippet.md", initSnippet)

	docs := filepath.Join(tmp, "docs")
	mdx := "{/* SDK_SNIPPET:RENDER:test-sdk/install hash=0 version=0.0.0 */}\n" +
		"```bash\n" +
		"placeholder\n" +
		"```\n\n" +
		"{/* SDK_SNIPPET:RENDER:test-sdk/init hash=0 version=0.0.0 */}\n" +
		"```python\n" +
		"placeholder\n" +
		"```\n"
	writeFile(t, docs, "page.mdx", mdx)

	changed, err := Render(os.DirFS(sdks), []string{docs})
	if err != nil {
		t.Fatal(err)
	}
	if len(changed) != 1 {
		t.Fatalf("expected 1 file changed, got %v", changed)
	}

	out, _ := os.ReadFile(filepath.Join(docs, "page.mdx"))
	got := string(out)

	// Install: empty `version` runtime-default → conditional omits → bare
	// `pip install foo` (no version pin).
	if !strings.Contains(got, "pip install foo\n") {
		t.Errorf("install body missing or wrong: \n%s", got)
	}
	if strings.Contains(got, "pip install foo==") {
		t.Errorf("install body should not have version pin (empty default): \n%s", got)
	}

	// Init: flag-key type substitutes YOUR_FLAG_KEY placeholder.
	if !strings.Contains(got, `foo.variation("YOUR_FLAG_KEY")`) {
		t.Errorf("init body should have YOUR_FLAG_KEY placeholder: \n%s", got)
	}

	// Hash and version stamps populated.
	if strings.Contains(got, "hash=0 ") {
		t.Errorf("placeholder hash should have been replaced: \n%s", got)
	}
	if strings.Contains(got, "version=0.0.0") {
		t.Errorf("placeholder version should have been replaced: \n%s", got)
	}

	// Verify passes on the rendered output.
	if err := Verify(os.DirFS(sdks), []string{docs}); err != nil {
		t.Errorf("verify after render should pass, got %v", err)
	}

	// Idempotent: a second render should produce no changes.
	changed, err = Render(os.DirFS(sdks), []string{docs})
	if err != nil {
		t.Fatal(err)
	}
	if len(changed) != 0 {
		t.Errorf("second render should be a no-op, got %v", changed)
	}
}

// Verify catches a hand-edit inside the fence body (hash mismatch) and
// fails loudly without rewriting anything.
func TestVerify_RejectsHandEdit(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	writeFile(t, sdks, "test-sdk/sdk.yaml", "id: test-sdk\n")
	writeFile(t, sdks, "test-sdk/snippets/install.snippet.md", installSnippet)

	docs := filepath.Join(tmp, "docs")
	writeFile(t, docs, "page.mdx",
		"{/* SDK_SNIPPET:RENDER:test-sdk/install hash=0 version=0.0.0 */}\n"+
			"```bash\n"+
			"placeholder\n"+
			"```\n")
	if _, err := Render(os.DirFS(sdks), []string{docs}); err != nil {
		t.Fatal(err)
	}

	// Hand-edit the rendered body.
	rendered, _ := os.ReadFile(filepath.Join(docs, "page.mdx"))
	doctored := strings.Replace(string(rendered), "pip install foo", "pip install something-else", 1)
	if doctored == string(rendered) {
		t.Fatal("doctoring failed; fixture out of sync")
	}
	os.WriteFile(filepath.Join(docs, "page.mdx"), []byte(doctored), 0o644)

	err := Verify(os.DirFS(sdks), []string{docs})
	if err == nil || !strings.Contains(err.Error(), "hand-edit detected") {
		t.Errorf("expected hand-edit error, got %v", err)
	}
}

// A surrounding <CodeBlock title='…'> wrapper is preserved across renders —
// only the fence body is owned by the snippet adapter.
func TestRender_PreservesCodeBlockWrapper(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	writeFile(t, sdks, "test-sdk/sdk.yaml", "id: test-sdk\n")
	writeFile(t, sdks, "test-sdk/snippets/install.snippet.md", installSnippet)

	docs := filepath.Join(tmp, "docs")
	mdx := "{/* SDK_SNIPPET:RENDER:test-sdk/install hash=0 version=0.0.0 */}\n" +
		"<CodeBlocks>\n" +
		"<CodeBlock title='Shell'>\n" +
		"\n" +
		"```bash\n" +
		"placeholder\n" +
		"```\n" +
		"\n" +
		"</CodeBlock>\n" +
		"</CodeBlocks>\n"
	writeFile(t, docs, "page.mdx", mdx)

	if _, err := Render(os.DirFS(sdks), []string{docs}); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(filepath.Join(docs, "page.mdx"))
	got := string(out)

	// The wrapper survives intact.
	if !strings.Contains(got, "<CodeBlock title='Shell'>") {
		t.Errorf("wrapper should be preserved: \n%s", got)
	}
	if !strings.Contains(got, "</CodeBlocks>") {
		t.Errorf("outer wrapper should be preserved: \n%s", got)
	}
	// Body filled in.
	if !strings.Contains(got, "pip install foo") {
		t.Errorf("body should be filled: \n%s", got)
	}
}
