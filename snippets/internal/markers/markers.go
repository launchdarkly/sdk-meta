package markers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// HashContent returns the first 12 hex chars of the SHA-256 of the bytes.
// This is the hash written into render markers.
func HashContent(content string) string {
	sum := sha256.Sum256([]byte(content))
	return hex.EncodeToString(sum[:])[:12]
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
	styleLine       commentStyle = iota // `// SDK_SNIPPET:...\n`
	styleJSXExpr                         // `{/* SDK_SNIPPET:... */}`
	styleBlock                           // `/* SDK_SNIPPET:... */`
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

// Hash returns the hash of the current content in the element's children.
func (m Match) Hash(src string) string {
	return HashContent(src[m.RegionStart:m.RegionEnd])
}

// Regex for the marker line inside any comment syntax. Captures fields.
var markerRe = regexp.MustCompile(
	`SDK_SNIPPET:RENDER:(?P<id>\S+)` +
		`(?:\s+hash=(?P<hash>[0-9a-fA-F]+))?` +
		`(?:\s+version=(?P<version>\S+))?` +
		`(?:\s+scope=(?P<scope>content|element|file))?`,
)

// tagNameRe matches the beginning of a JSX component tag: `<TagName`.
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
		// String literal: skip until matching close. Prevents false positives
		// inside quoted text like `"// SDK_SNIPPET:..."`.
		if src[i] == '"' || src[i] == '\'' || src[i] == '`' {
			quote := src[i]
			j := i + 1
			for j < len(src) {
				if src[j] == '\\' {
					j += 2
					continue
				}
				if src[j] == quote {
					j++
					break
				}
				j++
			}
			i = j
			continue
		}
		i++
	}
	return out, nil
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
		case '"', '\'', '`':
			quote := c
			j++
			for j < len(src) {
				if src[j] == '\\' {
					j += 2
					continue
				}
				if src[j] == quote {
					break
				}
				j++
			}
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

	// Find matching </TagName>. First-slice does not support same-tag nesting.
	closeTag := "</" + tag + ">"
	closeIdx := strings.Index(src[regionStart:], closeTag)
	if closeIdx < 0 {
		return Match{}, fmt.Errorf("marker %q: no </%s> closing tag found", m.ID, tag)
	}
	regionEnd := regionStart + closeIdx
	closeTagEnd := regionEnd + len(closeTag)

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
