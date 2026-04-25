package render

import (
	"fmt"
	"regexp"
	"strings"
)

// The snippet templating language is intentionally tiny:
//
//	{{ varName }}                substitute the value of an input
//	{{ if varName }}...{{ end }} emit "..." only if the input is truthy (non-empty)
//
// Conditionals do not nest in the first-pass slice. The inner "..." may still
// contain {{ varName }} substitutions. Future phases can extend this (filters,
// region toggles, version toggles) without breaking existing snippets.

type Node interface{ isNode() }

type Literal struct{ Text string }
type Var struct{ Name string }
type Cond struct {
	Var  string
	Body []Node
}

func (*Literal) isNode() {}
func (*Var) isNode()     {}
func (*Cond) isNode()    {}

var nameRe = `[a-zA-Z][a-zA-Z0-9_]*`
var tokenRe = regexp.MustCompile(`\{\{\s*(if\s+(` + nameRe + `)\s*|end\s*|(` + nameRe + `)\s*)\}\}`)

// Parse parses the mini-templating syntax into a flat node list.
// Conditionals are flattened: a Cond node contains its inner body.
func Parse(src string) ([]Node, error) {
	matches := tokenRe.FindAllStringSubmatchIndex(src, -1)

	var (
		i      = 0
		out    []Node
		stack  []*Cond
		append_ = func(n Node) {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				top.Body = append(top.Body, n)
			} else {
				out = append(out, n)
			}
		}
	)
	for _, m := range matches {
		start, end := m[0], m[1]
		if start > i {
			append_(&Literal{Text: src[i:start]})
		}
		token := src[start:end]
		switch {
		case m[4] >= 0: // "if NAME"
			name := src[m[4]:m[5]]
			c := &Cond{Var: name}
			stack = append(stack, c)
		case strings.HasPrefix(strings.TrimSpace(token[2:len(token)-2]), "end"):
			if len(stack) == 0 {
				return nil, fmt.Errorf("template: unmatched {{ end }} at offset %d", start)
			}
			closed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			append_(closed)
		case m[6] >= 0: // "NAME"
			name := src[m[6]:m[7]]
			append_(&Var{Name: name})
		default:
			return nil, fmt.Errorf("template: unrecognized directive %q at offset %d", token, start)
		}
		i = end
	}
	if i < len(src) {
		append_(&Literal{Text: src[i:]})
	}
	if len(stack) > 0 {
		return nil, fmt.Errorf("template: unclosed {{ if %s }}", stack[len(stack)-1].Var)
	}
	return out, nil
}
