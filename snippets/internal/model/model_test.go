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
	s, err := ParseFile(os.DirFS(tmp), filepath.Base(path))
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
	s, err := ParseFile(os.DirFS(tmp), filepath.Base(path))
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

// EffectiveChecks: legacy single-validator snippets synthesize one
// `kind: runtime` check from their flat fields.
func TestEffectiveChecks_LegacyFlatFields(t *testing.T) {
	v := Validation{
		Scaffold:     "x/y/z",
		Placeholders: map[string]string{"YOUR_KEY": "LAUNCHDARKLY_SDK_KEY"},
	}
	got := v.EffectiveChecks()
	if len(got) != 1 {
		t.Fatalf("want 1 synthesized check, got %d", len(got))
	}
	if got[0].Kind != "runtime" {
		t.Errorf("want kind=runtime, got %q", got[0].Kind)
	}
	if got[0].Scaffold != "x/y/z" {
		t.Errorf("want scaffold x/y/z, got %q", got[0].Scaffold)
	}
	if got[0].Placeholders["YOUR_KEY"] != "LAUNCHDARKLY_SDK_KEY" {
		t.Errorf("placeholders did not propagate")
	}
}

// EffectiveChecks: snippet with no validation fields returns nil.
func TestEffectiveChecks_Empty(t *testing.T) {
	if got := (Validation{}).EffectiveChecks(); got != nil {
		t.Errorf("want nil, got %+v", got)
	}
}

// EffectiveChecks: explicit checks inherit parent fields where unset
// and override where set.
func TestEffectiveChecks_ExplicitWithInheritance(t *testing.T) {
	v := Validation{
		Scaffold:     "parent/scaffold",
		Runtime:      "python",
		Placeholders: map[string]string{"YOUR_KEY": "LAUNCHDARKLY_SDK_KEY"},
		Env:          map[string]string{"PARENT_KEY": "parent-val"},
		Checks: []Check{
			{Kind: "parse"}, // inherits everything
			{
				Kind:     "runtime",
				Scaffold: "child/scaffold", // overrides parent
				Env:      map[string]string{"PARENT_KEY": "child-val", "CHILD_KEY": "child-only"},
			},
		},
	}
	got := v.EffectiveChecks()
	if len(got) != 2 {
		t.Fatalf("want 2 checks, got %d", len(got))
	}
	// First check inherits parent scaffold + runtime.
	if got[0].Scaffold != "parent/scaffold" || got[0].Runtime != "python" {
		t.Errorf("first check inheritance: scaffold=%q runtime=%q",
			got[0].Scaffold, got[0].Runtime)
	}
	if got[0].Env["PARENT_KEY"] != "parent-val" {
		t.Errorf("first check env inheritance failed: %+v", got[0].Env)
	}
	// Second check overrides scaffold + env value, inherits runtime.
	if got[1].Scaffold != "child/scaffold" {
		t.Errorf("second check override failed: scaffold=%q", got[1].Scaffold)
	}
	if got[1].Runtime != "python" {
		t.Errorf("second check should inherit runtime, got %q", got[1].Runtime)
	}
	if got[1].Env["PARENT_KEY"] != "child-val" {
		t.Errorf("second check env override failed: PARENT_KEY=%q", got[1].Env["PARENT_KEY"])
	}
	if got[1].Env["CHILD_KEY"] != "child-only" {
		t.Errorf("second check should retain its own env: %+v", got[1].Env)
	}
}
