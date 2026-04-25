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

// RenderForLDApplication produces the body that sits between the surrounding
// TSX backticks for the ld-application adapter. Substitutions become `${name}`
// expressions; conditionals become `${name ? `inner` : ''}` ternaries (nesting
// template literals, which is legal JS since the inner backticks are inside an
// ${} expression).
//
// The caller wraps the result in `...` when embedding in JSX.
func RenderForLDApplication(nodes []Node) string {
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
			sb.WriteString(RenderForLDApplication(x.Body))
			sb.WriteString("` : ''}")
		}
	}
	return sb.String()
}

// escapeTL escapes literal content for inclusion inside a JS template literal.
// Order matters: escape backslashes first, then backticks and ${ sequences.
func escapeTL(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, "`", "\\`")
	s = strings.ReplaceAll(s, "${", "\\${")
	return s
}
