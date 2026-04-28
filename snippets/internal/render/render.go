package render

import (
	"fmt"
	"strings"
)

// RenderRuntime substitutes inputs as concrete values. Used by the validator.
// Names that don't appear in the inputs map round-trip as literal `{{ name }}`
// so foreign template languages embedded in the snippet body (e.g. Vue's
// own `{{ ... }}` mustache syntax) survive validation untouched.
func RenderRuntime(nodes []Node, inputs map[string]string) (string, error) {
	var sb strings.Builder
	for _, n := range nodes {
		switch x := n.(type) {
		case *Literal:
			sb.WriteString(x.Text)
		case *Var:
			v, ok := inputs[x.Name]
			if !ok {
				// Unknown name — emit verbatim so foreign templates pass through.
				sb.WriteString("{{ ")
				sb.WriteString(x.Name)
				sb.WriteString(" }}")
				continue
			}
			sb.WriteString(v)
		case *Cond:
			v, ok := inputs[x.Var]
			if !ok {
				// We only support {{ if name }}…{{ end }} for declared inputs.
				// An undeclared name in a conditional is almost certainly an
				// authoring mistake (Vue's `v-if` is a different syntax).
				return "", fmt.Errorf("render: conditional refers to undeclared input %q", x.Var)
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

// HasInterpolation reports whether the template contains any Var (matching
// a declared input) or Cond node. Used by adapters to decide between bare-
// text and template-literal output forms. Foreign-template Vars (names
// not in `declaredInputs`) don't count — they're literal text.
func HasInterpolation(nodes []Node, declaredInputs map[string]struct{}) bool {
	for _, n := range nodes {
		switch x := n.(type) {
		case *Var:
			if _, ok := declaredInputs[x.Name]; ok {
				return true
			}
		case *Cond:
			return true
		}
	}
	return false
}

// RenderForLDApplicationTemplate produces the body for embedding inside a
// JS template literal: substitutions become `${name}`, conditionals become
// `${name ? `inner` : ''}`. Literal text is escaped so backslashes,
// backticks, and `${` sequences in the source survive into the runtime
// string verbatim. Names not in declaredInputs are emitted as literal
// `{{ name }}` so foreign-template syntax (Vue's mustaches) passes through.
// The caller wraps the returned bytes in backticks.
func RenderForLDApplicationTemplate(nodes []Node, declaredInputs map[string]struct{}) string {
	var sb strings.Builder
	for _, n := range nodes {
		switch x := n.(type) {
		case *Literal:
			sb.WriteString(escapeTL(x.Text))
		case *Var:
			if _, ok := declaredInputs[x.Name]; !ok {
				// Foreign template — emit the original `{{ name }}` literally,
				// escaped for the surrounding template literal.
				sb.WriteString(escapeTL("{{ " + x.Name + " }}"))
				continue
			}
			sb.WriteString("${")
			sb.WriteString(x.Name)
			sb.WriteString("}")
		case *Cond:
			sb.WriteString("${")
			sb.WriteString(x.Var)
			sb.WriteString(" ? `")
			sb.WriteString(RenderForLDApplicationTemplate(x.Body, declaredInputs))
			sb.WriteString("` : ''}")
		}
	}
	return sb.String()
}

// RenderForJSXText produces the body for embedding as bare JSX text (i.e.
// not wrapped in a template literal). Only valid when the template has no
// interpolation matching a declared input. Foreign-template `{{ name }}`
// passes through verbatim.
//
// Backslashes, backticks, and ${} sequences are emitted verbatim — JSX
// text does not interpret any of those. Literal `{` and `}` characters
// are not handled here; the caller should fall back to template-literal
// output if either appears.
func RenderForJSXText(nodes []Node, declaredInputs map[string]struct{}) (string, error) {
	var sb strings.Builder
	for _, n := range nodes {
		switch x := n.(type) {
		case *Literal:
			sb.WriteString(x.Text)
		case *Var:
			if _, ok := declaredInputs[x.Name]; ok {
				return "", fmt.Errorf("RenderForJSXText: template has declared interpolation; use RenderForLDApplicationTemplate")
			}
			// Foreign template — emit literal.
			sb.WriteString("{{ ")
			sb.WriteString(x.Name)
			sb.WriteString(" }}")
		case *Cond:
			return "", fmt.Errorf("RenderForJSXText: template has conditional; use RenderForLDApplicationTemplate")
		}
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
