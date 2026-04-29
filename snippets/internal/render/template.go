package render

import (
	"fmt"
	"regexp"
	"strings"
)

// The snippet templating language is intentionally tiny:
//
//	{{ varName }}                substitute the value of an input
//	{{ varName | filter }}       substitute the value with a filter applied
//	{{ if varName }}...{{ end }} emit "..." only if the input is truthy (non-empty)
//
// Conditionals do not nest in the first-pass slice. The inner "..." may still
// contain {{ varName }} substitutions. Filters supported today: `camelCase`.
// React's snippets need the camelCase filter because `useFlags()` destructures
// camelCased identifiers; ld-application maps these to `${camelCase(name)}`.

type Node interface{ isNode() }

type Literal struct{ Text string }
type Var struct {
	Name   string
	Filter string // empty if no filter; e.g. "camelCase"
}
type Cond struct {
	Var  string
	Body []Node
}

func (*Literal) isNode() {}
func (*Var) isNode()     {}
func (*Cond) isNode()    {}

var nameRe = `[a-zA-Z][a-zA-Z0-9_]*`
var tokenRe = regexp.MustCompile(`\{\{\s*(if\s+(` + nameRe + `)\s*|end\s*|(` + nameRe + `)(?:\s*\|\s*(` + nameRe + `))?\s*)\}\}`)

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
		// Note: order of cases matters. "if NAME" must be tested first so that a
		// hypothetical variable starting with "if" can never be reached. Equality
		// is required for the "end" case so a variable like {{ endTime }} doesn't
		// get treated as a block-close.
		switch {
		case m[4] >= 0: // "if NAME"
			name := src[m[4]:m[5]]
			c := &Cond{Var: name}
			stack = append(stack, c)
		case strings.TrimSpace(token[2:len(token)-2]) == "end":
			if len(stack) == 0 {
				return nil, fmt.Errorf("template: unmatched {{ end }} at offset %d", start)
			}
			closed := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			append_(closed)
		case m[6] >= 0: // "NAME" optionally followed by "| FILTER"
			name := src[m[6]:m[7]]
			v := &Var{Name: name}
			if m[8] >= 0 {
				filter := src[m[8]:m[9]]
				if filter != "camelCase" {
					return nil, fmt.Errorf("template: unknown filter %q at offset %d (only `camelCase` is supported)", filter, start)
				}
				v.Filter = filter
			}
			append_(v)
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
