package model

import (
	"os"
	"path/filepath"
	"testing"
)

// Regression: a blank line immediately following the opening fence used
// to be silently dropped from CodeBody because the open-fence regex's
// trailing `\s*` (greedy, includes \n) consumed the line terminator,
// and firstCodeBlock then unconditionally skipped one more newline.
// The fix uses `[ \t]*` so only horizontal whitespace can follow the
// language tag.
func TestParseFile_PreservesLeadingBlankLine(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "x.snippet.md")
	body := "---\n" +
		"id: x/y\n" +
		"sdk: x\n" +
		"kind: install\n" +
		"lang: bash\n" +
		"---\n\n" +
		"Description text.\n\n" +
		"```bash\n" +
		"\n" +
		"first real line\n" +
		"last line\n" +
		"```\n"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := ParseFile(path)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	want := "\nfirst real line\nlast line"
	if s.CodeBody != want {
		t.Fatalf("CodeBody mismatch\n got:  %q\n want: %q", s.CodeBody, want)
	}
}

// Confirms the no-blank-line case still works as before — the trailing
// newline before the closing fence is stripped, and the body is the raw
// code with no extra leading whitespace.
func TestParseFile_NoLeadingBlankLine(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "x.snippet.md")
	body := "---\n" +
		"id: x/y\n" +
		"sdk: x\n" +
		"kind: hello-world\n" +
		"lang: python\n" +
		"---\n\n" +
		"```python\n" +
		"print('hi')\n" +
		"```\n"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	s, err := ParseFile(path)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if s.CodeBody != "print('hi')" {
		t.Fatalf("CodeBody = %q", s.CodeBody)
	}
	if s.CodeLang != "python" {
		t.Fatalf("CodeLang = %q", s.CodeLang)
	}
}
