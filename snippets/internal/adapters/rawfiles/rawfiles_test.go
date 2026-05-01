package rawfiles

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeSDK seeds a one-snippet sdks/<id>/ tree at sdksDir. The body is
// plain text so the no-input render path is exercised.
func writeSDK(t *testing.T, sdksDir, sdkID, snippetID, lang, body string) {
	t.Helper()
	d := filepath.Join(sdksDir, sdkID)
	if err := os.MkdirAll(filepath.Join(d, "snippets", "sdk-info"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(d, "sdk.yaml"),
		[]byte("id: "+sdkID+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	md := "---\n" +
		"id: " + snippetID + "\n" +
		"sdk: " + sdkID + "\n" +
		"kind: install\n" +
		"lang: " + lang + "\n" +
		"file: out/" + filepath.Base(snippetID) + ".txt\n" +
		"---\n\n" +
		"```" + lang + "\n" + body + "\n```\n"
	if err := os.WriteFile(filepath.Join(d, "snippets", "sdk-info", filepath.Base(snippetID)+".snippet.md"),
		[]byte(md), 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeFile(t *testing.T, p, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// Round-trip: a manifest pointing at a registered snippet writes the body
// (plus a trailing newline) to the configured output path.
func TestRender_WritesSnippetBodiesToManifestPaths(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	consumer := filepath.Join(tmp, "consumer")
	writeSDK(t, sdks, "x-sdk", "x-sdk/sdk-info/install", "shell", "npm i x")

	manifest := filepath.Join(consumer, "extract.yaml")
	writeFile(t, manifest, "out: snippets\nfiles:\n  - id: x-sdk/sdk-info/install\n    path: x-sdk/install.txt\n")

	written, err := Render(os.DirFS(sdks), manifest, "")
	if err != nil {
		t.Fatal(err)
	}
	if len(written) != 1 {
		t.Fatalf("expected 1 file written, got %d: %v", len(written), written)
	}
	got, err := os.ReadFile(filepath.Join(consumer, "snippets", "x-sdk", "install.txt"))
	if err != nil {
		t.Fatal(err)
	}
	// The .snippet.md fence syntax strips the trailing newline before the
	// closing fence; the adapter adds one back so the on-disk file ends
	// with a newline (matches the source .txt POSIX convention).
	if string(got) != "npm i x\n" {
		t.Fatalf("body mismatch: %q", got)
	}
}

// The --consumer override resolves the manifest's `out:` against the
// supplied root rather than the manifest's directory. This is the path
// gonfalon's CI uses to render into a checkout that's separate from the
// repo where the manifest lives.
func TestRender_ConsumerOverrideResolvesOut(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	manifestDir := filepath.Join(tmp, "manifests")
	consumer := filepath.Join(tmp, "elsewhere")
	writeSDK(t, sdks, "x-sdk", "x-sdk/sdk-info/install", "shell", "npm i x")

	manifest := filepath.Join(manifestDir, "extract.yaml")
	writeFile(t, manifest, "out: snippets\nfiles:\n  - id: x-sdk/sdk-info/install\n    path: a.txt\n")

	if _, err := Render(os.DirFS(sdks), manifest, consumer); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(consumer, "snippets", "a.txt")); err != nil {
		t.Fatalf("expected file under consumer/snippets: %v", err)
	}
	if _, err := os.Stat(filepath.Join(manifestDir, "snippets", "a.txt")); !os.IsNotExist(err) {
		t.Fatalf("expected NO file under manifest dir, got err=%v", err)
	}
}

// A manifest referencing an unknown snippet ID fails with a useful
// suggestions list rather than silently emitting an empty file.
func TestRender_RejectsUnknownID(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	consumer := filepath.Join(tmp, "consumer")
	writeSDK(t, sdks, "x-sdk", "x-sdk/sdk-info/install", "shell", "npm i x")

	manifest := filepath.Join(consumer, "extract.yaml")
	writeFile(t, manifest, "out: snippets\nfiles:\n  - id: x-sdk/sdk-info/innstall\n    path: x.txt\n")

	_, err := Render(os.DirFS(sdks), manifest, "")
	if err == nil {
		t.Fatal("expected error for unknown id")
	}
	if !strings.Contains(err.Error(), "unknown snippet id") {
		t.Fatalf("unexpected error: %v", err)
	}
}

// Manifest validation rejects duplicate IDs (would silently shadow a
// previous render entry).
func TestLoadManifest_RejectsDuplicateID(t *testing.T) {
	tmp := t.TempDir()
	manifest := filepath.Join(tmp, "extract.yaml")
	writeFile(t, manifest, "out: out\nfiles:\n  - id: a/sdk-info/x\n    path: a.txt\n  - id: a/sdk-info/x\n    path: b.txt\n")
	if _, _, err := LoadManifest(manifest); err == nil ||
		!strings.Contains(err.Error(), "duplicate id") {
		t.Fatalf("want duplicate-id error, got %v", err)
	}
}

// Manifest validation rejects duplicate output paths (two snippets racing
// to the same file).
func TestLoadManifest_RejectsDuplicatePath(t *testing.T) {
	tmp := t.TempDir()
	manifest := filepath.Join(tmp, "extract.yaml")
	writeFile(t, manifest, "out: out\nfiles:\n  - id: a/sdk-info/x\n    path: a.txt\n  - id: a/sdk-info/y\n    path: a.txt\n")
	if _, _, err := LoadManifest(manifest); err == nil ||
		!strings.Contains(err.Error(), "duplicate output path") {
		t.Fatalf("want duplicate-path error, got %v", err)
	}
}

// Manifest validation rejects entries whose `path:` would write outside
// the declared `out:` root.
func TestLoadManifest_RejectsTraversingPath(t *testing.T) {
	tmp := t.TempDir()
	manifest := filepath.Join(tmp, "extract.yaml")
	writeFile(t, manifest, "out: snippets\nfiles:\n  - id: a/sdk-info/x\n    path: ../escape.txt\n")
	if _, _, err := LoadManifest(manifest); err == nil ||
		!strings.Contains(err.Error(), "must not escape") {
		t.Fatalf("want escape error, got %v", err)
	}
}

// An unknown frontmatter key in the manifest is rejected at parse time
// (KnownFields is on). Catches typos like `output:` instead of `out:`.
func TestLoadManifest_RejectsUnknownKey(t *testing.T) {
	tmp := t.TempDir()
	manifest := filepath.Join(tmp, "extract.yaml")
	writeFile(t, manifest, "out: out\noutput: something\nfiles:\n  - id: a/sdk-info/x\n    path: a.txt\n")
	if _, _, err := LoadManifest(manifest); err == nil ||
		!strings.Contains(err.Error(), "field output") {
		t.Fatalf("want unknown-field error, got %v", err)
	}
}

// Missing `out:` is rejected — the renderer wouldn't know where to write.
func TestLoadManifest_RejectsMissingOut(t *testing.T) {
	tmp := t.TempDir()
	manifest := filepath.Join(tmp, "extract.yaml")
	writeFile(t, manifest, "files:\n  - id: a/sdk-info/x\n    path: a.txt\n")
	if _, _, err := LoadManifest(manifest); err == nil ||
		!strings.Contains(err.Error(), "out:") {
		t.Fatalf("want missing-out error, got %v", err)
	}
}

// atomicWriteFile preserves the destination's permission bits when
// overwriting. Mirrors the contract of ldapplication.atomicWriteFile.
func TestAtomicWriteFile_PreservesMode(t *testing.T) {
	tmp := t.TempDir()
	dest := filepath.Join(tmp, "f.txt")
	if err := os.WriteFile(dest, []byte("first"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := atomicWriteFile(dest, []byte("second")); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(dest)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("expected mode 0600 preserved, got %o", info.Mode().Perm())
	}
	got, _ := os.ReadFile(dest)
	if string(got) != "second" {
		t.Fatalf("body mismatch: %q", got)
	}
}

// atomicWriteFile creates intermediate directories when called from
// Render — exercises that the Render path mkdir-p's.
func TestRender_CreatesIntermediateDirs(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	consumer := filepath.Join(tmp, "consumer")
	writeSDK(t, sdks, "x-sdk", "x-sdk/sdk-info/install", "shell", "npm i x")

	manifest := filepath.Join(consumer, "extract.yaml")
	writeFile(t, manifest, "out: out\nfiles:\n  - id: x-sdk/sdk-info/install\n    path: deep/nested/path/install.txt\n")

	if _, err := Render(os.DirFS(sdks), manifest, ""); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(consumer, "out", "deep", "nested", "path", "install.txt")); err != nil {
		t.Fatalf("expected nested file: %v", err)
	}
}

// Bodies that have a foreign-template `{{ name }}` (not declared as an
// input) round-trip verbatim through the renderer when called with no
// inputs. Important for the cursor prompt's `{{SDK_NAME}}` placeholders
// — they should land in the rendered file unchanged so gonfalon's
// runtime substitution sees the same source it always has.
//
// Caveat documented in the inconsistency log: `{{SDK_NAME}}` (no inner
// whitespace) round-trips as `{{ SDK_NAME }}` (with whitespace) because
// the renderer normalizes Var formatting.
func TestRender_PreservesForeignTemplateMarkers(t *testing.T) {
	tmp := t.TempDir()
	sdks := filepath.Join(tmp, "sdks")
	consumer := filepath.Join(tmp, "consumer")
	// Body holds a `{{ NAME }}` that's not in any inputs map.
	writeSDK(t, sdks, "x-sdk", "x-sdk/sdk-info/install", "text", "Hello {{ NAME }}")

	manifest := filepath.Join(consumer, "extract.yaml")
	writeFile(t, manifest, "out: out\nfiles:\n  - id: x-sdk/sdk-info/install\n    path: hello.txt\n")

	if _, err := Render(os.DirFS(sdks), manifest, ""); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(consumer, "out", "hello.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "Hello {{ NAME }}\n" {
		t.Fatalf("expected foreign-template preserved, got %q", got)
	}
}
