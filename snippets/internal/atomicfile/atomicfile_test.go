package atomicfile

import (
	"os"
	"path/filepath"
	"testing"
)

// Write preserves the destination's permission bits when overwriting,
// so a checkout that has tightened a file (e.g. read-only mode for a
// CODEOWNER-protected file) doesn't quietly get reset to 0644.
func TestWrite_PreservesMode(t *testing.T) {
	tmp := t.TempDir()
	dest := filepath.Join(tmp, "f.txt")
	if err := os.WriteFile(dest, []byte("first"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := Write(dest, []byte("second")); err != nil {
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

// Write to a non-existing destination uses 0644 as the default mode.
func TestWrite_NewFileDefaultMode(t *testing.T) {
	tmp := t.TempDir()
	dest := filepath.Join(tmp, "new.txt")
	if err := Write(dest, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(dest)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o644 {
		t.Fatalf("expected mode 0644 for new file, got %o", info.Mode().Perm())
	}
}

// Write leaves no .tmp turds behind in the destination directory after
// a successful write.
func TestWrite_NoTempfileLeftover(t *testing.T) {
	tmp := t.TempDir()
	dest := filepath.Join(tmp, "f.txt")
	if err := Write(dest, []byte("body")); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != "f.txt" {
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		t.Fatalf("expected only f.txt in dir, got %v", names)
	}
}
