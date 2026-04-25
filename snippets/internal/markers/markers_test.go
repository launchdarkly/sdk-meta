package markers

import "testing"

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
