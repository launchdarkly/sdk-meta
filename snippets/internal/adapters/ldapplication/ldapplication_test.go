package ldapplication

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Helper: write a tiny sdks/ tree pointing at a fixture TSX file.
func writeSDKTree(t *testing.T, sdksDir, sdkID, getStartedRel, appDir string) {
	t.Helper()
	d := filepath.Join(sdksDir, sdkID)
	if err := os.MkdirAll(filepath.Join(d, "snippets", "getting-started"), 0o755); err != nil {
		t.Fatal(err)
	}
	yaml := "id: " + sdkID + "\n" +
		"sdk-meta-id: x\n" +
		"display-name: X\n" +
		"type: server-side\n" +
		"languages:\n  - id: x\n    extensions: [\".x\"]\n" +
		"package-managers: [pip]\n" +
		"regions: [commercial]\n" +
		"hello-world-repo: x/y\n" +
		"ld-application:\n  get-started-file: " + getStartedRel + "\n" +
		"docs:\n  reference-page: /x\n"
	if err := os.WriteFile(filepath.Join(d, "sdk.yaml"), []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}
}

// Regression for review #3: a sdk.yaml that points get-started-file outside
// of appDir must be rejected, not followed.
func TestDiscoverTargetFiles_RejectsTraversal(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	writeSDKTree(t, sdks, "evil-sdk", "../escape.tsx", app)

	_, err := discoverTargetFiles(os.DirFS(sdks), app)
	if err == nil || !strings.Contains(err.Error(), "escapes appDir") {
		t.Fatalf("want escapes-appDir error, got %v", err)
	}
}

func TestDiscoverTargetFiles_RejectsAbsolute(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	writeSDKTree(t, sdks, "evil-sdk", "/etc/passwd", app)

	_, err := discoverTargetFiles(os.DirFS(sdks), app)
	if err == nil || !strings.Contains(err.Error(), "must be relative") {
		t.Fatalf("want must-be-relative error, got %v", err)
	}
}

// Per the scope=content contract, attributes are the consumer's to choose.
// `verify` must NOT reject an attribute-only edit — only changes to the
// element's children should fail. Tests below exercise both cases.
func TestVerify_AcceptsAttributeEdit(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: shell
---

`+"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	tsx := `// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0
<Snippet lang="shell">
  placeholder
</Snippet>
`
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(tsx), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	// Hand-edit an attribute (add withCopyButton). Verify must still pass.
	bytes, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	edited := strings.Replace(string(bytes),
		`<Snippet lang="shell">`,
		`<Snippet lang="shell" withCopyButton>`, 1)
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatalf("verify should accept attribute-only edit: %v", err)
	}
}

func TestVerify_RejectsChildEdit(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: shell
---

`+"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	tsx := `// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0
<Snippet lang="shell">
  placeholder
</Snippet>
`
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(tsx), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	bytes, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	edited := strings.Replace(string(bytes), "mkdir hi", "mkdir HACKED", 1)
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
	err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), app)
	if err == nil || !strings.Contains(err.Error(), "hand-edit detected") {
		t.Fatalf("verify should reject child edit, got %v", err)
	}
}

func TestNeedsTemplateLiteral_BareTextStaysBare(t *testing.T) {
	// A snippet with no interpolation, no newline, no JSX-special chars
	// should render bare so the existing in-place style is preserved.
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: shell
---

`+"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	tsx := `// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0
<Snippet lang="shell">
  placeholder
</Snippet>
`
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(tsx), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	// The body must be the bare text, NOT wrapped in {`...`}.
	if !strings.Contains(string(out), "<Snippet lang=\"shell\">\n  mkdir hi\n</Snippet>") {
		t.Fatalf("expected bare-text rendering, got:\n%s", out)
	}
	// And verify must accept it.
	if err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatalf("verify after render should pass: %v", err)
	}
}

// Regression: a snippet whose only interpolation is a foreign-template
// `{{ name }}` (e.g. Vue's mustache) must render via the template-
// literal path so the curlies get escaped. Bare-text rendering would
// emit `{{ name }}` straight into JSX, where the outer `{` opens a JS
// expression and the inner becomes an object-literal — invalid JSX.
func TestNeedsTemplateLiteral_ForeignTemplateUsesTemplateLiteral(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Snippet body has no declared inputs and no JSX-special chars in
	// any literal — just a foreign-template `{{ flagValue }}`.
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: html
---

`+"```html\nflag is {{ flagValue }}\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	tsx := `// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0
<Snippet lang="html">
  placeholder
</Snippet>
`
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(tsx), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	// The body MUST be wrapped in a template literal: the foreign-
	// template `{{ flagValue }}` would otherwise be interpreted by JSX.
	if !strings.Contains(string(out), "{`") {
		t.Fatalf("expected template-literal wrapping, got:\n%s", out)
	}
	if !strings.Contains(string(out), "{{ flagValue }}") {
		t.Fatalf("expected foreign-template literal preserved, got:\n%s", out)
	}
	// Verify must accept what render produced.
	if err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatalf("verify after render should pass: %v", err)
	}
}

// Render preserves the existing `version=` field on a marker when the
// rendered body is byte-identical to what's already on disk. The version
// is meant to record the binary that last *changed* this snippet's
// content, not the binary that last touched the file — otherwise every
// release would rewrite every marker in every consumer file even when
// no snippet actually changed.
func TestRender_PreservesVersionWhenBodyUnchanged(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: shell
---

`+"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}

	// First render: marker has hash=0 so the body has to be filled in;
	// the binary's version stamp lands.
	initial := `// SDK_SNIPPET:RENDER:x/cmd hash=0 version=0.0.1-old
<Snippet lang="shell">
  placeholder
</Snippet>
`
	tsx := filepath.Join(app, "app.tsx")
	if err := os.WriteFile(tsx, []byte(initial), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(tsx)
	// The marker version should now be the binary's current version —
	// the old `0.0.1-old` is overwritten because the body is changing
	// (placeholder → real content).
	if strings.Contains(string(out), "version=0.0.1-old") {
		t.Fatalf("first render should have replaced stale version stamp; got:\n%s", out)
	}

	// Second render: body is identical now, so the marker's version stays
	// where it landed — even if we hand-edit the version on disk to a
	// non-current value before re-rendering.
	current, _ := os.ReadFile(tsx)
	doctored := strings.Replace(string(current), "version=", "version=9.9.9-pinned-", 1)
	if doctored == string(current) {
		t.Fatal("doctoring failed; test fixture out of sync")
	}
	if err := os.WriteFile(tsx, []byte(doctored), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), app); err != nil {
		t.Fatal(err)
	}
	out, _ = os.ReadFile(tsx)
	if !strings.Contains(string(out), "version=9.9.9-pinned-") {
		t.Fatalf("second render should have preserved the pinned version; got:\n%s", out)
	}
}

// Regression for review #5: a marker with no hash= field must be rejected
// during verify.
func TestVerify_RejectsMissingHash(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks", "x")
	if err := os.MkdirAll(filepath.Join(sdks, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"), []byte(
		"id: x\nsdk-meta-id: y\ndisplay-name: y\ntype: server-side\n"+
			"languages:\n  - id: y\n    extensions: [\".y\"]\n"+
			"package-managers: []\nregions: []\nhello-world-repo: a/b\n"+
			"ld-application:\n  get-started-file: app.tsx\n"+
			"docs:\n  reference-page: /\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		`---
id: x/cmd
sdk: x
kind: bootstrap
lang: shell
---

`+"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	if err := os.MkdirAll(app, 0o755); err != nil {
		t.Fatal(err)
	}
	tsx := `// SDK_SNIPPET:RENDER:x/cmd version=0.1.0
<Snippet lang="shell">
  mkdir hi
</Snippet>
`
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(tsx), 0o644); err != nil {
		t.Fatal(err)
	}
	err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), app)
	if err == nil || !strings.Contains(err.Error(), "missing required hash") {
		t.Fatalf("want missing-hash error, got %v", err)
	}
}
