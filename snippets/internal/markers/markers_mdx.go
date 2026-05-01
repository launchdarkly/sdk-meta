package markers

import (
	"fmt"
	"regexp"
	"strings"
)

// MDXMatch describes one render marker in an MDX file. Unlike the TSX
// adapter, the marker doesn't anchor to a JSX element — it anchors to the
// next fenced code block. The hash covers the bytes BETWEEN the fences,
// so the consumer is free to wrap the fence in <CodeBlocks>, set the
// `title=` attribute on a surrounding <CodeBlock>, etc., without a
// re-render. This matches the scope=content contract.
type MDXMatch struct {
	Fields MarkerFields

	// CommentStart..CommentEnd is the marker comment itself (the {/* ... */} bytes).
	CommentStart, CommentEnd int

	// FenceStart is the byte index of the first backtick of the opening
	// fence line. FenceEnd is one byte past the closing fence's final
	// newline (or end-of-file if the fence closes the file).
	FenceStart, FenceEnd int

	// RegionStart..RegionEnd is the fence body — bytes after the opening
	// fence's terminating newline, up to (and excluding) the closing fence
	// line's leading newline. The hash is computed over this slice.
	RegionStart, RegionEnd int

	// Lang is the info string immediately after the opening backticks
	// (e.g., "python", "bash", or empty for an untagged fence).
	Lang string
}

// fenceLineRe matches a fenced-code-block delimiter line. The CommonMark
// spec allows up to three leading spaces and three or more backticks; we
// allow that, capturing the backtick run and the optional info string.
// Closing fences must use the same backtick count as the opener — the
// scanner enforces that downstream.
var fenceLineRe = regexp.MustCompile(`^([ ]{0,3})(` + "`" + `{3,})([^\n` + "`" + `]*)$`)

// ScanMDX finds every {/* SDK_SNIPPET:RENDER:... */} comment in an MDX file
// and pairs each one with the next fenced code block. Wrapping JSX elements
// (<CodeBlocks>, <CodeBlock title='…'>) between the comment and the fence
// are skipped — only their attributes/content are attribution-checked, so
// the consumer can choose those freely.
//
// Returns matches in file order. If a marker is followed by something that
// isn't a fenced code block (within ~50 lines), it's an error — markers
// must be physically adjacent to the code they own.
func ScanMDX(src string) ([]MDXMatch, error) {
	var out []MDXMatch
	i := 0
	for i < len(src) {
		// MDX prose is mostly markdown; the only comment syntax we care
		// about is the JSX-expression `{/* ... */}` form. HTML comments
		// `<!-- ... -->` are intentionally NOT recognized — pick one and
		// stick to it.
		if i+2 < len(src) && src[i] == '{' && src[i+1] == '/' && src[i+2] == '*' {
			rel := strings.Index(src[i+3:], "*/}")
			if rel < 0 {
				return nil, fmt.Errorf("unterminated {/* */} comment at offset %d", i)
			}
			inner := src[i+3 : i+3+rel]
			closeEnd := i + 3 + rel + 3
			if m, ok := parseMarker(inner); ok {
				match, err := attachFence(src, i, closeEnd, m)
				if err != nil {
					return nil, err
				}
				out = append(out, match)
				i = match.FenceEnd
				continue
			}
			i = closeEnd
			continue
		}
		i++
	}
	return out, nil
}

// attachFence locates the next fenced code block after the marker comment
// and packages the bytes into an MDXMatch. Wrapping JSX elements between
// the comment and the fence are tolerated — only their absence of a fence
// is an error. The "next fence" is the first opening fence within
// maxLinesBeforeFence lines of the comment; further away suggests an
// authoring mistake (marker drifted away from its code block) and is
// rejected.
func attachFence(src string, commentStart, commentEnd int, m MarkerFields) (MDXMatch, error) {
	if m.Scope != "content" {
		return MDXMatch{}, fmt.Errorf("marker %q: scope=%q is not supported in lddocs adapter", m.ID, m.Scope)
	}

	// Walk lines from commentEnd looking for the opening fence. Cap the
	// search so a forgotten fence doesn't silently consume hundreds of
	// lines of unrelated prose.
	const maxLinesBeforeFence = 50
	lineStart := commentEnd
	// Skip the rest of the line containing the comment close (the comment
	// itself isn't on its own line necessarily — but in practice it almost
	// always is). Move lineStart to the next line's first byte.
	if nl := strings.IndexByte(src[lineStart:], '\n'); nl >= 0 {
		lineStart += nl + 1
	} else {
		return MDXMatch{}, fmt.Errorf("marker %q: comment is on the last line; no fence follows", m.ID)
	}

	openFenceStart := -1
	openFenceLineEnd := -1
	openTicks := ""
	openLang := ""
	linesScanned := 0
	for lineStart < len(src) && linesScanned < maxLinesBeforeFence {
		lineEnd := strings.IndexByte(src[lineStart:], '\n')
		var line string
		var lineEndAbs int
		if lineEnd < 0 {
			line = src[lineStart:]
			lineEndAbs = len(src)
		} else {
			line = src[lineStart : lineStart+lineEnd]
			lineEndAbs = lineStart + lineEnd
		}
		if sub := fenceLineRe.FindStringSubmatch(line); sub != nil {
			openTicks = sub[2]
			openLang = strings.TrimSpace(sub[3])
			openFenceStart = lineStart + len(sub[1])
			openFenceLineEnd = lineEndAbs
			break
		}
		// Anything else (blank lines, JSX wrapper tags, prose) is fine —
		// keep scanning.
		linesScanned++
		if lineEnd < 0 {
			lineStart = len(src)
			break
		}
		lineStart = lineEndAbs + 1
	}
	if openFenceStart < 0 {
		return MDXMatch{}, fmt.Errorf("marker %q: no fenced code block found within %d lines after comment", m.ID, maxLinesBeforeFence)
	}

	// Body starts on the line after the opening fence.
	if openFenceLineEnd >= len(src) {
		return MDXMatch{}, fmt.Errorf("marker %q: opening fence is at end of file", m.ID)
	}
	regionStart := openFenceLineEnd + 1

	// Find the matching closing fence: the first subsequent line whose
	// (trimmed-leading-whitespace) prefix is exactly the same backtick run
	// followed by optional whitespace.
	closeRe := regexp.MustCompile(`(?m)^[ ]{0,3}` + openTicks + `[ \t]*$`)
	closeIdx := closeRe.FindStringIndex(src[regionStart:])
	if closeIdx == nil {
		return MDXMatch{}, fmt.Errorf("marker %q: unterminated fenced code block (opener %q)", m.ID, openTicks)
	}
	closeAbs := regionStart + closeIdx[0]
	// Region ends at the byte BEFORE the closing fence line's leading
	// newline. closeAbs points at the first byte of the closing fence
	// line; the newline that ends the previous line is at closeAbs-1.
	// Empty body (opener and closer on consecutive lines) clamps to regionStart.
	regionEnd := max(closeAbs-1, regionStart)
	fenceEnd := regionStart + closeIdx[1]
	// Consume the trailing newline after the closing fence if present, so
	// FenceEnd points at the next line's first byte (matches FenceStart's
	// "first byte of the fence line" convention).
	if fenceEnd < len(src) && src[fenceEnd] == '\n' {
		fenceEnd++
	}

	return MDXMatch{
		Fields:       m,
		CommentStart: commentStart,
		CommentEnd:   commentEnd,
		FenceStart:   openFenceStart,
		FenceEnd:     fenceEnd,
		RegionStart:  regionStart,
		RegionEnd:    regionEnd,
		Lang:         openLang,
	}, nil
}

// FormatMDXMarker renders the marker comment for MDX (always JSX-expression
// style — line/block comments don't apply outside JSX expression context).
func FormatMDXMarker(f MarkerFields) string {
	return FormatMarker(styleJSXExpr, f)
}
