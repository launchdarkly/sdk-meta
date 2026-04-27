package markers

import (
	"strings"
	"testing"
)

func TestScanTSX_LineCommentAndJSXComment(t *testing.T) {
	src := `import { Snippet } from 'x';

const A = () => (
  // SDK_SNIPPET:RENDER:foo/bar hash=abc version=0.1.0
  <Snippet lang="python">{` + "`" + `hello ${x}` + "`" + `}</Snippet>
);

const B = () => (
  <li>
    {/* SDK_SNIPPET:RENDER:foo/baz hash=def version=0.1.0 */}
    <Snippet lang="shell">mkdir hello</Snippet>
  </li>
);
`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
	if matches[0].Fields.ID != "foo/bar" || matches[0].TagName != "Snippet" {
		t.Fatalf("match 0: %+v", matches[0])
	}
	if matches[0].Style != styleLine {
		t.Fatalf("match 0 style: want styleLine, got %v", matches[0].Style)
	}
	body0 := src[matches[0].RegionStart:matches[0].RegionEnd]
	if body0 != "{`hello ${x}`}" {
		t.Fatalf("body 0: %q", body0)
	}
	if matches[1].Fields.ID != "foo/baz" || matches[1].Style != styleJSXExpr {
		t.Fatalf("match 1: %+v", matches[1])
	}
	body1 := src[matches[1].RegionStart:matches[1].RegionEnd]
	if body1 != "mkdir hello" {
		t.Fatalf("body 1: %q", body1)
	}
}

func TestScanTSX_IgnoresMarkerInsideString(t *testing.T) {
	src := `const s = "// SDK_SNIPPET:RENDER:nope hash=000 version=0.1.0";
`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 0 {
		t.Fatalf("expected 0 matches, got %d", len(matches))
	}
}

func TestScanTSX_NoMarkers(t *testing.T) {
	src := `import x from 'y';\nconst A = () => <div>hi</div>;\n`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 0 {
		t.Fatalf("expected 0 matches, got %d", len(matches))
	}
}

func TestScanTSX_BlockCommentMarker(t *testing.T) {
	src := `const A = (
  /* SDK_SNIPPET:RENDER:foo/bar hash=abc version=0.1.0 */
  <Snippet>body</Snippet>
);`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 1 || matches[0].Style != styleBlock {
		t.Fatalf("want one block-style match, got %+v", matches)
	}
}

func TestScanTSX_UnterminatedJSXComment(t *testing.T) {
	src := `<div>{/* SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0`
	_, err := ScanTSX(src)
	if err == nil {
		t.Fatal("want error for unterminated comment")
	}
}

func TestScanTSX_NonWhitespaceBetweenMarkerAndElement(t *testing.T) {
	src := `// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0
const x = 1;
<Snippet>hi</Snippet>`
	_, err := ScanTSX(src)
	if err == nil || !strings.Contains(err.Error(), "non-whitespace") {
		t.Fatalf("want non-whitespace error, got %v", err)
	}
}

func TestScanTSX_SelfClosingElement(t *testing.T) {
	src := `// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0
<Snippet />`
	_, err := ScanTSX(src)
	if err == nil || !strings.Contains(err.Error(), "self-closing") {
		t.Fatalf("want self-closing error, got %v", err)
	}
}

func TestScanTSX_MissingClosingTag(t *testing.T) {
	src := `// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0
<Snippet>only opening`
	_, err := ScanTSX(src)
	if err == nil || !strings.Contains(err.Error(), "no matching") {
		t.Fatalf("want missing-close error, got %v", err)
	}
}

// Regression for review #6: nested same-tag should track depth so the
// outer </Snippet> is matched, not the first inner one.
func TestScanTSX_NestedSameTag(t *testing.T) {
	src := `// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0
<Snippet lang="outer"><Snippet lang="inner">inner</Snippet>after</Snippet>`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	body := src[matches[0].RegionStart:matches[0].RegionEnd]
	want := `<Snippet lang="inner">inner</Snippet>after`
	if body != want {
		t.Fatalf("body mismatch:\n got:  %q\n want: %q", body, want)
	}
}

// `<SnippetGroup` should not count as a `<Snippet` open at depth.
func TestScanTSX_SimilarPrefixTag(t *testing.T) {
	src := `// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0
<Snippet><SnippetGroup>x</SnippetGroup></Snippet>`
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if len(matches) != 1 {
		t.Fatalf("want 1 match, got %d", len(matches))
	}
	body := src[matches[0].RegionStart:matches[0].RegionEnd]
	if body != `<SnippetGroup>x</SnippetGroup>` {
		t.Fatalf("body mismatch: %q", body)
	}
}

// Regression for review #9: a nested template literal inside a ${ } expression
// inside a backtick string must not end the outer string scan early.
func TestScanTSX_BacktickWithNestedTemplate(t *testing.T) {
	src := "const x = `outer ${fn(`inner` + 1)} done`;\n" +
		"// SDK_SNIPPET:RENDER:foo hash=abc version=0.1.0\n" +
		"<Snippet>body</Snippet>"
	matches, err := ScanTSX(src)
	if err != nil {
		t.Fatalf("scan: %v", err)
	}
	if len(matches) != 1 || matches[0].Fields.ID != "foo" {
		t.Fatalf("want one match for foo, got %+v", matches)
	}
}

func TestParseMarker_OptionalFields(t *testing.T) {
	// hash, version, scope all absent: parseMarker still extracts the ID.
	got, ok := parseMarker(" SDK_SNIPPET:RENDER:foo/bar")
	if !ok || got.ID != "foo/bar" {
		t.Fatalf("parse: %+v ok=%v", got, ok)
	}
	if got.Hash != "" || got.Version != "" {
		t.Fatalf("expected empty optional fields: %+v", got)
	}
	if got.Scope != "content" {
		t.Fatalf("default scope should be content: %q", got.Scope)
	}
}
