package markers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// hashLen is the number of hex characters of the SHA-256 written into render
// markers. Twelve hex chars is ~48 bits — enough collision resistance to
// catch accidental drift in CI but NOT a cryptographic integrity claim. Do
// not extend other security-sensitive checks to rely on this prefix without
// widening it first.
const hashLen = 12

// HashContent returns the first hashLen hex chars of the SHA-256 of the bytes.
// This is the hash written into render markers.
func HashContent(content string) string {
	sum := sha256.Sum256([]byte(content))
	return hex.EncodeToString(sum[:])[:hashLen]
}

// MarkerFields is the set of key=value pairs we recognize after the ID.
type MarkerFields struct {
	ID      string
	Hash    string
	Version string
	Scope   string // "content" (default), "element", or "file"
}

// commentStyle records how the marker was spelled so a rewrite preserves it.
type commentStyle int

const (
	styleLine    commentStyle = iota // `// SDK_SNIPPET:...\n`
	styleJSXExpr                     // `{/* SDK_SNIPPET:... */}`
	styleBlock                       // `/* SDK_SNIPPET:... */`
)

// Match describes one marker occurrence in a TSX/TS file.
type Match struct {
	Fields MarkerFields
	Style  commentStyle

	// Byte offsets: [CommentStart, CommentEnd) is the marker comment itself.
	CommentStart, CommentEnd int

	// The JSX element that follows. For scope=content:
	//   [RegionStart, RegionEnd) is the element's children (between > and </).
	// The opening tag and attributes remain at [OpenTagStart, RegionStart).
	OpenTagStart, RegionStart, RegionEnd, CloseTagEnd int

	TagName string
}

// Regex for the marker line inside any comment syntax. Captures fields.
// Only ID is required; hash/version/scope are optional in the regex but the
// hash is REQUIRED at verify time (see ldapplication.Verify).
var markerRe = regexp.MustCompile(
	`SDK_SNIPPET:RENDER:(?P<id>\S+)` +
		`(?:\s+hash=(?P<hash>[0-9a-fA-F]+))?` +
		`(?:\s+version=(?P<version>\S+))?` +
		`(?:\s+scope=(?P<scope>content|element|file))?`,
)

// tagNameRe matches the beginning of a JSX component tag: `<TagName`.
//
// Constraint: the leading character is uppercase. JSX semantics distinguish
// component identifiers (capitalized — resolve to a value in scope) from DOM
// elements (lowercase — emitted as raw HTML). Snippet authors who need to
// mark a DOM element directly will need to wrap it in a component first.
var tagNameRe = regexp.MustCompile(`<([A-Z][A-Za-z0-9]*)`)

// ScanTSX finds every render marker in a TSX file and pairs it with the
// following JSX element. Returns matches in file order.
func ScanTSX(src string) ([]Match, error) {
	var out []Match

	// Walk byte by byte; detect each of the three comment syntaxes and pass
	// the comment content through the marker regex. This keeps things simple
	// without a full TSX parser.
	i := 0
	for i < len(src) {
		// Line comment: `//` to newline
		if i+1 < len(src) && src[i] == '/' && src[i+1] == '/' {
			end := strings.Index(src[i:], "\n")
			if end < 0 {
				end = len(src) - i
			}
			inner := src[i+2 : i+end]
			if m, ok := parseMarker(inner); ok {
				match, err := attachElement(src, i, i+end, styleLine, m)
				if err != nil {
					return nil, err
				}
				out = append(out, match)
				i = match.CloseTagEnd
				continue
			}
			i += end
			continue
		}
		// JSX expression comment: `{/*` ... `*/}`
		if i+2 < len(src) && src[i] == '{' && src[i+1] == '/' && src[i+2] == '*' {
			end := strings.Index(src[i:], "*/}")
			if end < 0 {
				return nil, fmt.Errorf("unterminated {/* */} comment at offset %d", i)
			}
			inner := src[i+3 : i+end]
			if m, ok := parseMarker(inner); ok {
				match, err := attachElement(src, i, i+end+3, styleJSXExpr, m)
				if err != nil {
					return nil, err
				}
				out = append(out, match)
				i = match.CloseTagEnd
				continue
			}
			i += end + 3
			continue
		}
		// Block comment: `/*` ... `*/`
		if i+1 < len(src) && src[i] == '/' && src[i+1] == '*' {
			end := strings.Index(src[i:], "*/")
			if end < 0 {
				return nil, fmt.Errorf("unterminated /* */ comment at offset %d", i)
			}
			inner := src[i+2 : i+end]
			if m, ok := parseMarker(inner); ok {
				match, err := attachElement(src, i, i+end+2, styleBlock, m)
				if err != nil {
					return nil, err
				}
				out = append(out, match)
				i = match.CloseTagEnd
				continue
			}
			i += end + 2
			continue
		}
		// String / template literal: skip until matching close. Prevents false
		// positives inside quoted text like `"// SDK_SNIPPET:..."` and inside
		// backtick template literals (which can themselves contain ${} expressions
		// containing nested template literals).
		if src[i] == '"' || src[i] == '\'' {
			i = skipPlainString(src, i)
			continue
		}
		if src[i] == '`' {
			i = skipBacktick(src, i)
			continue
		}
		i++
	}
	return out, nil
}

// skipPlainString consumes a "..." or '...' string literal starting at i and
// returns the offset immediately after the closing quote. Backslash escapes
// are honored. If the string is unterminated the function returns len(src).
func skipPlainString(src string, i int) int {
	quote := src[i]
	j := i + 1
	for j < len(src) {
		if src[j] == '\\' {
			j += 2
			continue
		}
		if src[j] == quote {
			return j + 1
		}
		j++
	}
	return len(src)
}

// skipBacktick consumes a `...` template literal starting at i and returns
// the offset immediately after the closing backtick. Inside the literal,
// `${ ... }` expressions are recognized and their bodies are scanned with
// brace-balanced depth so a nested template literal does not prematurely
// close the outer one.
func skipBacktick(src string, i int) int {
	j := i + 1
	for j < len(src) {
		switch src[j] {
		case '\\':
			j += 2
		case '`':
			return j + 1
		case '$':
			if j+1 < len(src) && src[j+1] == '{' {
				j = skipJSExpr(src, j+1) // pass the `{`
				continue
			}
			j++
		default:
			j++
		}
	}
	return len(src)
}

// skipJSExpr consumes a `{ ... }` JS expression starting at the byte indexed
// by `i` (which must point at the opening `{`) and returns the offset
// immediately after the matching `}`. Strings, template literals, and nested
// `{}` blocks are tracked. Comments inside the expression are NOT recognized
// (they don't appear in JSX-attribute / template-literal contexts in
// practice); add support if a real snippet host requires it.
func skipJSExpr(src string, i int) int {
	depth := 0
	j := i
	for j < len(src) {
		switch src[j] {
		case '{':
			depth++
			j++
		case '}':
			depth--
			j++
			if depth == 0 {
				return j
			}
		case '"', '\'':
			j = skipPlainString(src, j)
		case '`':
			j = skipBacktick(src, j)
		case '\\':
			j += 2
		default:
			j++
		}
	}
	return len(src)
}

func parseMarker(inner string) (MarkerFields, bool) {
	sub := markerRe.FindStringSubmatch(inner)
	if sub == nil {
		return MarkerFields{}, false
	}
	get := func(name string) string {
		idx := markerRe.SubexpIndex(name)
		if idx < 0 {
			return ""
		}
		return sub[idx]
	}
	f := MarkerFields{
		ID:      get("id"),
		Hash:    get("hash"),
		Version: get("version"),
		Scope:   get("scope"),
	}
	if f.Scope == "" {
		f.Scope = "content"
	}
	return f, true
}

// attachElement finds the JSX element that the marker introduces.
// First-slice scope is "content" only; element/file come later.
func attachElement(src string, commentStart, commentEnd int, style commentStyle, m MarkerFields) (Match, error) {
	if m.Scope != "content" {
		return Match{}, fmt.Errorf("marker %q: scope=%q is not supported in first-pass implementation", m.ID, m.Scope)
	}
	// Find `<TagName` after the comment, skipping whitespace.
	rest := src[commentEnd:]
	offs := tagNameRe.FindStringSubmatchIndex(rest)
	if offs == nil {
		return Match{}, fmt.Errorf("marker %q: no JSX element found after comment", m.ID)
	}
	// Allow only whitespace between the comment and `<`.
	gap := rest[:offs[0]]
	if strings.TrimSpace(gap) != "" {
		return Match{}, fmt.Errorf("marker %q: non-whitespace between comment and element: %q", m.ID, gap)
	}
	tag := rest[offs[2]:offs[3]]
	openTagStart := commentEnd + offs[0]

	// Find end of opening tag: the `>` that closes `<TagName ...>`. Track
	// strings and brace depth so `>` inside attributes (e.g., `onClick={x > 0}`)
	// doesn't fool us.
	depth := 0
	j := commentEnd + offs[1] // right after the tag name
	for j < len(src) {
		c := src[j]
		if depth == 0 {
			if c == '/' && j+1 < len(src) && src[j+1] == '>' {
				return Match{}, fmt.Errorf("marker %q: element <%s/> is self-closing — scope=content requires a body", m.ID, tag)
			}
			if c == '>' {
				break
			}
		}
		switch c {
		case '"', '\'':
			j = skipPlainString(src, j)
			continue
		case '`':
			j = skipBacktick(src, j)
			continue
		case '{':
			depth++
		case '}':
			depth--
		}
		j++
	}
	if j >= len(src) {
		return Match{}, fmt.Errorf("marker %q: unterminated <%s opening tag", m.ID, tag)
	}
	regionStart := j + 1 // byte after `>`

	// Find matching </TagName> at the same nesting depth. Walk the body,
	// counting <TagName ...> opens and </TagName> closes; the close that
	// brings depth back to zero is the match. String/backtick/JSX-expression
	// scanning ensures angle brackets inside attributes or string literals
	// don't perturb the depth.
	openMarker := "<" + tag
	closeMarker := "</" + tag + ">"
	tagDepth := 1
	regionEnd := -1
	closeTagEnd := -1
	k := regionStart
	for k < len(src) {
		switch src[k] {
		case '"', '\'':
			k = skipPlainString(src, k)
			continue
		case '`':
			k = skipBacktick(src, k)
			continue
		case '<':
			// Same-tag open at depth: only if followed by a non-identifier byte
			// (so `<Snippet` matches but `<SnippetGroup` doesn't).
			if hasTagPrefix(src, k, openMarker) {
				tagDepth++
				k += len(openMarker)
				continue
			}
			if strings.HasPrefix(src[k:], closeMarker) {
				tagDepth--
				if tagDepth == 0 {
					regionEnd = k
					closeTagEnd = k + len(closeMarker)
					k = closeTagEnd
					goto done
				}
				k += len(closeMarker)
				continue
			}
		}
		k++
	}
done:
	if regionEnd < 0 {
		return Match{}, fmt.Errorf("marker %q: no matching </%s> closing tag found", m.ID, tag)
	}

	return Match{
		Fields:       m,
		Style:        style,
		CommentStart: commentStart,
		CommentEnd:   commentEnd,
		OpenTagStart: openTagStart,
		RegionStart:  regionStart,
		RegionEnd:    regionEnd,
		CloseTagEnd:  closeTagEnd,
		TagName:      tag,
	}, nil
}

// hasTagPrefix reports whether src[i:] starts with prefix AND the byte
// immediately after is one that can't be part of a longer identifier.
// Used to distinguish `<Snippet` from `<SnippetGroup`.
func hasTagPrefix(src string, i int, prefix string) bool {
	if !strings.HasPrefix(src[i:], prefix) {
		return false
	}
	end := i + len(prefix)
	if end >= len(src) {
		return true
	}
	c := src[end]
	if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return false
	}
	return true
}

// FormatMarker renders the marker comment in the given style, with hash+version.
func FormatMarker(style commentStyle, f MarkerFields) string {
	body := fmt.Sprintf("SDK_SNIPPET:RENDER:%s hash=%s version=%s", f.ID, f.Hash, f.Version)
	if f.Scope != "" && f.Scope != "content" {
		body += " scope=" + f.Scope
	}
	switch style {
	case styleLine:
		return "// " + body
	case styleJSXExpr:
		return "{/* " + body + " */}"
	case styleBlock:
		return "/* " + body + " */"
	}
	return "// " + body
}
