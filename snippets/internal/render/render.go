package render

import (
	"fmt"
	"strings"
)

// RenderRuntime substitutes inputs as concrete values. Used by the validator.
// A missing input is an error — validation must always have a value.
func RenderRuntime(nodes []Node, inputs map[string]string) (string, error) {
	var sb strings.Builder
	for _, n := range nodes {
		switch x := n.(type) {
		case *Literal:
			sb.WriteString(x.Text)
		case *Var:
			v, ok := inputs[x.Name]
			if !ok {
				return "", fmt.Errorf("render: missing runtime input %q", x.Name)
			}
			sb.WriteString(v)
		case *Cond:
			v, ok := inputs[x.Var]
			if !ok {
				return "", fmt.Errorf("render: missing runtime input %q (used in conditional)", x.Var)
			}
			if v != "" {
				inner, err := RenderRuntime(x.Body, inputs)
				if err != nil {
					return "", err
				}
				sb.WriteString(inner)
			}
		}
	}
	return sb.String(), nil
}

// HasInterpolation reports whether the template contains any Var or Cond
// node anywhere in the tree. Used by adapters to decide between bare-text
// and template-literal output forms.
func HasInterpolation(nodes []Node) bool {
	for _, n := range nodes {
		switch x := n.(type) {
		case *Var:
			return true
		case *Cond:
			return true
		case *Literal:
			_ = x
		}
	}
	return false
}

// RenderForLDApplicationTemplate produces the body for embedding inside a
// JS template literal: substitutions become `${name}`, conditionals become
// `${name ? `inner` : ''}`. Literal text is escaped so backslashes,
// backticks, and `${` sequences in the source survive into the runtime
// string verbatim. The caller wraps the returned bytes in backticks.
func RenderForLDApplicationTemplate(nodes []Node) string {
	var sb strings.Builder
	for _, n := range nodes {
		switch x := n.(type) {
		case *Literal:
			sb.WriteString(escapeTL(x.Text))
		case *Var:
			sb.WriteString("${")
			sb.WriteString(x.Name)
			sb.WriteString("}")
		case *Cond:
			sb.WriteString("${")
			sb.WriteString(x.Var)
			sb.WriteString(" ? `")
			sb.WriteString(RenderForLDApplicationTemplate(x.Body))
			sb.WriteString("` : ''}")
		}
	}
	return sb.String()
}

// RenderForJSXText produces the body for embedding as bare JSX text (i.e.
// not wrapped in a template literal). Only valid when the template has no
// interpolation; the caller is expected to consult HasInterpolation first.
//
// Backslashes, backticks, and ${} sequences are emitted verbatim — JSX
// text does not interpret any of those. Literal `{` and `}` characters
// are not handled here (they would require JSX-specific escaping); the
// caller should fall back to template-literal output if either appears.
func RenderForJSXText(nodes []Node) (string, error) {
	var sb strings.Builder
	for _, n := range nodes {
		l, ok := n.(*Literal)
		if !ok {
			return "", fmt.Errorf("RenderForJSXText: template has interpolation; use RenderForLDApplicationTemplate")
		}
		sb.WriteString(l.Text)
	}
	return sb.String(), nil
}

// ContainsJSXSpecial reports whether the rendered bare text contains
// characters that JSX would interpret specially (`{` or `}`). When true,
// the caller should switch to template-literal output.
func ContainsJSXSpecial(s string) bool {
	return strings.ContainsAny(s, "{}")
}

// escapeTL escapes literal content for inclusion inside a JS template literal.
// Order matters: escape backslashes first, then backticks and ${ sequences.
func escapeTL(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, "`", "\\`")
	s = strings.ReplaceAll(s, "${", "\\${")
	return s
}
