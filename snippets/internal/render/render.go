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
				sb.WriteString(literalVar(x))
				continue
			}
			if x.Filter != "" {
				v = applyFilter(x.Filter, v)
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
				sb.WriteString(escapeTL(literalVar(x)))
				continue
			}
			sb.WriteString("${")
			if x.Filter != "" {
				sb.WriteString(x.Filter)
				sb.WriteString("(")
				sb.WriteString(x.Name)
				sb.WriteString(")")
			} else {
				sb.WriteString(x.Name)
			}
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
			sb.WriteString(literalVar(x))
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

// literalVar formats a Var node back to its source `{{ name }}` /
// `{{ name | filter }}` form. Used when emitting an undeclared name
// verbatim so foreign-template syntax round-trips intact.
func literalVar(v *Var) string {
	if v.Filter == "" {
		return "{{ " + v.Name + " }}"
	}
	return "{{ " + v.Name + " | " + v.Filter + " }}"
}

// applyFilter applies a filter to a runtime value. Today only `camelCase`
// is supported — used by react-client-sdk's snippets where useFlags()
// destructures camelCased identifiers from a kebab-cased flag key.
func applyFilter(name, value string) string {
	switch name {
	case "camelCase":
		return camelCase(value)
	default:
		return value
	}
}

// camelCase mirrors @gonfalon/strings' camelCase. Converts kebab-case,
// snake_case, and space-separated words to camelCase. Leading non-alpha
// runs are stripped. The first segment stays lowercase; subsequent
// segments get an uppercase initial.
//
// Examples:
//   sample-feature  -> sampleFeature
//   my_flag_key     -> myFlagKey
//   already-camelOK -> alreadyCamelOk (lowercases later segments first)
func camelCase(s string) string {
	var segs []string
	var cur strings.Builder
	flush := func() {
		if cur.Len() > 0 {
			segs = append(segs, strings.ToLower(cur.String()))
			cur.Reset()
		}
	}
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			cur.WriteRune(r)
			continue
		}
		flush()
	}
	flush()
	if len(segs) == 0 {
		return ""
	}
	var out strings.Builder
	out.WriteString(segs[0])
	for _, seg := range segs[1:] {
		if seg == "" {
			continue
		}
		out.WriteString(strings.ToUpper(seg[:1]))
		out.WriteString(seg[1:])
	}
	return out.String()
}
