package ldapplication

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeSDK seeds a one-snippet sdks/<id>/ tree at sdksDir. The snippet's
// fenced body is plain shell so the bare-text rendering path is exercised.
func writeSDK(t *testing.T, sdksDir, sdkID string) {
	t.Helper()
	d := filepath.Join(sdksDir, sdkID)
	if err := os.MkdirAll(filepath.Join(d, "snippets"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d, "sdk.yaml"),
		[]byte("id: "+sdkID+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d, "snippets", "x.snippet.md"), []byte(
		"---\n"+
			"id: "+sdkID+"/cmd\n"+
			"file: app.tsx\n"+
			"---\n\n"+
			"```shell\nmkdir hi\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

// writeAppFile writes a TSX entry at <appDir>/<rel> with the given body.
func writeAppFile(t *testing.T, appDir, rel, body string) {
	t.Helper()
	p := filepath.Join(appDir, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// discoverFilesUnder skips junk dirs (node_modules etc.) and only picks
// up files containing the SDK_SNIPPET:RENDER sentinel.
func TestDiscoverFilesUnder_FiltersByExtensionAndSentinel(t *testing.T) {
	tmp := t.TempDir()
	app := filepath.Join(tmp, "app")
	writeAppFile(t, app, "with-marker.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=0 version=0.1.0\n<Snippet>x</Snippet>\n")
	writeAppFile(t, app, "no-marker.tsx", "export default function App() { return null }\n")
	writeAppFile(t, app, "wrong-ext.go", "// SDK_SNIPPET:RENDER:x/cmd hash=0\n")
	// node_modules should be skipped wholesale.
	writeAppFile(t, app, "node_modules/foo/bar.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=0\n<Snippet>x</Snippet>\n")

	files, err := discoverFilesUnder([]string{app})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || filepath.Base(files[0]) != "with-marker.tsx" {
		t.Fatalf("expected exactly with-marker.tsx, got %v", files)
	}
}

// Multiple entrypoints accumulate without duplicating files when the
// entrypoints overlap.
func TestDiscoverFilesUnder_DedupsOverlappingEntrypoints(t *testing.T) {
	tmp := t.TempDir()
	app := filepath.Join(tmp, "app")
	writeAppFile(t, app, "sub/marker.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=0 version=0.1.0\n<Snippet>x</Snippet>\n")

	files, err := discoverFilesUnder([]string{app, filepath.Join(app, "sub")})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file after dedup, got %d: %v", len(files), files)
	}
}

// The skip-list (node_modules, build, dist, …) prunes noise *under* the
// entrypoint root. If the user passes one of those names *as* the
// entrypoint (e.g. a project that emits its TSX into ./build), the walk
// must still descend — otherwise discovery silently produces zero files.
func TestDiscoverFilesUnder_RootMatchingSkipNameStillDescends(t *testing.T) {
	tmp := t.TempDir()
	root := filepath.Join(tmp, "build")
	writeAppFile(t, root, "marker.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=0 version=0.1.0\n<Snippet>x</Snippet>\n")

	files, err := discoverFilesUnder([]string{root})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 || filepath.Base(files[0]) != "marker.tsx" {
		t.Fatalf("expected marker.tsx under skip-named root, got %v", files)
	}
}

// A non-directory entrypoint is rejected loudly rather than silently no-op'd.
func TestDiscoverFilesUnder_RejectsNonDirectory(t *testing.T) {
	tmp := t.TempDir()
	f := filepath.Join(tmp, "not-a-dir.tsx")
	writeAppFile(t, tmp, "not-a-dir.tsx", "// SDK_SNIPPET:RENDER:x/cmd hash=0\n")
	if _, err := discoverFilesUnder([]string{f}); err == nil ||
		!strings.Contains(err.Error(), "not a directory") {
		t.Fatalf("expected not-a-directory error, got %v", err)
	}
}

// Per the scope=content contract, attributes are the consumer's to choose.
// `verify` must NOT reject an attribute-only edit — only changes to the
// element's children should fail. Tests below exercise both cases.
func TestVerify_AcceptsAttributeEdit(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	writeSDK(t, sdks, "x")
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0\n"+
			"<Snippet lang=\"shell\">\n  placeholder\n</Snippet>\n")

	if _, err := Render(os.DirFS(sdks), []string{app}); err != nil {
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
	if err := Verify(os.DirFS(sdks), []string{app}); err != nil {
		t.Fatalf("verify should accept attribute-only edit: %v", err)
	}
}

func TestVerify_RejectsChildEdit(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	writeSDK(t, sdks, "x")
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0\n"+
			"<Snippet lang=\"shell\">\n  placeholder\n</Snippet>\n")

	if _, err := Render(os.DirFS(sdks), []string{app}); err != nil {
		t.Fatal(err)
	}
	bytes, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	edited := strings.Replace(string(bytes), "mkdir hi", "mkdir HACKED", 1)
	if err := os.WriteFile(filepath.Join(app, "app.tsx"), []byte(edited), 0o644); err != nil {
		t.Fatal(err)
	}
	err := Verify(os.DirFS(sdks), []string{app})
	if err == nil || !strings.Contains(err.Error(), "hand-edit detected") {
		t.Fatalf("verify should reject child edit, got %v", err)
	}
}

func TestNeedsTemplateLiteral_BareTextStaysBare(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	writeSDK(t, sdks, "x")
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0\n"+
			"<Snippet lang=\"shell\">\n  placeholder\n</Snippet>\n")

	if _, err := Render(os.DirFS(sdks), []string{app}); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	if !strings.Contains(string(out), "<Snippet lang=\"shell\">\n  mkdir hi\n</Snippet>") {
		t.Fatalf("expected bare-text rendering, got:\n%s", out)
	}
	if err := Verify(os.DirFS(sdks), []string{app}); err != nil {
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
	if err := os.WriteFile(filepath.Join(sdks, "sdk.yaml"),
		[]byte("id: x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	// Snippet body has no declared inputs and no JSX-special chars in
	// any literal — just a foreign-template `{{ flagValue }}`.
	if err := os.WriteFile(filepath.Join(sdks, "snippets", "x.snippet.md"), []byte(
		"---\n"+
			"id: x/cmd\n"+
			"file: app.tsx\n"+
			"---\n\n"+
			"```html\nflag is {{ flagValue }}\n```\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	app := filepath.Join(tmp, "app")
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=000000000000 version=0.1.0\n"+
			"<Snippet lang=\"html\">\n  placeholder\n</Snippet>\n")

	if _, err := Render(os.DirFS(filepath.Join(tmp, "sdks")), []string{app}); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(filepath.Join(app, "app.tsx"))
	if !strings.Contains(string(out), "{`") {
		t.Fatalf("expected template-literal wrapping, got:\n%s", out)
	}
	if !strings.Contains(string(out), "{{ flagValue }}") {
		t.Fatalf("expected foreign-template literal preserved, got:\n%s", out)
	}
	if err := Verify(os.DirFS(filepath.Join(tmp, "sdks")), []string{app}); err != nil {
		t.Fatalf("verify after render should pass: %v", err)
	}
}

// Render preserves the existing `version=` field on a marker when the
// rendered body is byte-identical to what's already on disk. The version
// is meant to record the binary that last *changed* this snippet's
// content, not the binary that last touched the file.
func TestRender_PreservesVersionWhenBodyUnchanged(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	writeSDK(t, sdks, "x")
	tsx := filepath.Join(app, "app.tsx")
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd hash=0 version=0.0.1-old\n"+
			"<Snippet lang=\"shell\">\n  placeholder\n</Snippet>\n")

	if _, err := Render(os.DirFS(sdks), []string{app}); err != nil {
		t.Fatal(err)
	}
	out, _ := os.ReadFile(tsx)
	if strings.Contains(string(out), "version=0.0.1-old") {
		t.Fatalf("first render should have replaced stale version stamp; got:\n%s", out)
	}

	current, _ := os.ReadFile(tsx)
	doctored := strings.Replace(string(current), "version=", "version=9.9.9-pinned-", 1)
	if doctored == string(current) {
		t.Fatal("doctoring failed; test fixture out of sync")
	}
	if err := os.WriteFile(tsx, []byte(doctored), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Render(os.DirFS(sdks), []string{app}); err != nil {
		t.Fatal(err)
	}
	out, _ = os.ReadFile(tsx)
	if !strings.Contains(string(out), "version=9.9.9-pinned-") {
		t.Fatalf("second render should have preserved the pinned version; got:\n%s", out)
	}
}

// A marker with no hash= field must be rejected during verify.
func TestVerify_RejectsMissingHash(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	app := filepath.Join(tmp, "app")
	writeSDK(t, sdks, "x")
	// Note: no hash= field on the marker.
	writeAppFile(t, app, "app.tsx",
		"// SDK_SNIPPET:RENDER:x/cmd version=0.1.0\n"+
			"<Snippet lang=\"shell\">\n  mkdir hi\n</Snippet>\n")

	err := Verify(os.DirFS(sdks), []string{app})
	if err == nil || !strings.Contains(err.Error(), "missing required hash") {
		t.Fatalf("want missing-hash error, got %v", err)
	}
}
