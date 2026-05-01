package markers

import (
	"strings"
	"testing"
)

// Bare fence directly after the comment.
func TestScanMDX_BareFence(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo/bar hash=abc version=0.1.0 */}\n" +
		"```python\n" +
		"print('hello')\n" +
		"```\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	m := matches[0]
	if m.Fields.ID != "foo/bar" {
		t.Errorf("ID: %q", m.Fields.ID)
	}
	if m.Lang != "python" {
		t.Errorf("Lang: %q", m.Lang)
	}
	body := src[m.RegionStart:m.RegionEnd]
	if body != "print('hello')" {
		t.Errorf("body: %q", body)
	}
}

// CodeBlock JSX wrapper between the marker and the fence is fine.
func TestScanMDX_SkipsJSXWrappers(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo/bar hash=0 version=0.0.0 */}\n" +
		"<CodeBlocks>\n" +
		"<CodeBlock title='Python'>\n" +
		"\n" +
		"```python\n" +
		"print('hi')\n" +
		"```\n" +
		"\n" +
		"</CodeBlock>\n" +
		"</CodeBlocks>\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	body := src[matches[0].RegionStart:matches[0].RegionEnd]
	if body != "print('hi')" {
		t.Errorf("body: %q", body)
	}
}

// Multiple markers in one file each anchor to the next fence.
func TestScanMDX_MultipleMarkers(t *testing.T) {
	src := "## Install\n" +
		"{/* SDK_SNIPPET:RENDER:a hash=0 version=0 */}\n" +
		"```bash\n" +
		"pip install foo\n" +
		"```\n" +
		"\n" +
		"## Init\n" +
		"{/* SDK_SNIPPET:RENDER:b hash=0 version=0 */}\n" +
		"```python\n" +
		"import foo\n" +
		"```\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("want 2 matches, got %d", len(matches))
	}
	if matches[0].Fields.ID != "a" || matches[1].Fields.ID != "b" {
		t.Errorf("ids: %q, %q", matches[0].Fields.ID, matches[1].Fields.ID)
	}
	if matches[0].Lang != "bash" || matches[1].Lang != "python" {
		t.Errorf("langs: %q, %q", matches[0].Lang, matches[1].Lang)
	}
}

// Fence body containing backticks (e.g., shell command substitution) is
// preserved as long as the closing fence uses the same number of opening
// backticks at the start of its line. A 4-backtick opener tolerates a
// 3-backtick code line.
func TestScanMDX_FourBacktickFence(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo hash=0 version=0 */}\n" +
		"````markdown\n" +
		"```python\n" +
		"print('inside')\n" +
		"```\n" +
		"````\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	body := src[matches[0].RegionStart:matches[0].RegionEnd]
	want := "```python\nprint('inside')\n```"
	if body != want {
		t.Errorf("body:\n%q\nwant:\n%q", body, want)
	}
}

// A marker with no fence within the cap is rejected.
func TestScanMDX_NoFenceFound(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo hash=0 version=0 */}\n" +
		strings.Repeat("just prose, no fence\n", 60)
	if _, err := ScanMDX(src); err == nil {
		t.Fatal("want error, got nil")
	}
}

// An unterminated comment is rejected loudly rather than silently skipped.
func TestScanMDX_UnterminatedComment(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo hash=0 version=0\n```python\nx\n```\n"
	if _, err := ScanMDX(src); err == nil ||
		!strings.Contains(err.Error(), "unterminated") {
		t.Fatalf("want unterminated comment error, got %v", err)
	}
}

// Markers inside <CodeBlock> attributes (rare but possible) shouldn't
// break the scanner. The marker comment is at the prose level; whatever
// is inside the code block body is treated as opaque text.
func TestScanMDX_MarkerLikeStringInsideFence(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo hash=0 version=0 */}\n" +
		"```python\n" +
		"# {/* SDK_SNIPPET:RENDER:not-a-real-marker hash=0 version=0 */}\n" +
		"print('x')\n" +
		"```\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("want exactly 1 match (the comment-string inside the fence body must not be treated as a marker), got %d", len(matches))
	}
	if matches[0].Fields.ID != "foo" {
		t.Errorf("got id %q, want foo", matches[0].Fields.ID)
	}
}

// Empty fence body (consecutive opener/closer lines) clamps cleanly
// rather than producing a negative slice.
func TestScanMDX_EmptyBody(t *testing.T) {
	src := "{/* SDK_SNIPPET:RENDER:foo hash=0 version=0 */}\n" +
		"```\n" +
		"```\n"
	matches, err := ScanMDX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	body := src[matches[0].RegionStart:matches[0].RegionEnd]
	if body != "" {
		t.Errorf("body: %q want empty", body)
	}
}
