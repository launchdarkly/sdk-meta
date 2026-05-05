---
id: go-server-sdk/scaffolds/go-syntax-only
sdk: go-server-sdk
kind: scaffold
lang: go
file: main.go
description: |
  Parse-only validator for Go server SDK doc fragments.

  Go is unusually strict about where syntax may appear: top-level
  `import (...)` blocks and top-level `func` declarations cannot live
  inside another function body, so the simple "wrap body in
  `_wrappee()`" pattern breaks for snippets that show install-time
  imports or middleware-style helper functions.

  Instead of building the body, the harness uses `go/parser.ParseFile`
  on a synthetic file that splices the wrappee body either as a
  top-level fragment (when it begins with `import`/`func`/`package`/
  `var`/`const`/`type`) or wrapped inside a no-op function body
  otherwise. ParseFile walks the AST without resolving symbols, so
  unresolved package names (`ld`, `ldobserve`, etc.) don't fail the
  check.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed by go/parser.
validation:
  runtime: go
  entrypoint: main.go
---

```go
package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

const wrappeeBody = `
{{ body }}
`

// splitTopLevel separates a Go fragment into top-level declarations
// (`import (...)`, `func ...`, `var/const/type/package`) followed by
// optional function-body residue. The docs sometimes show both —
// "add this import AND use this code" — so the harness can't pick one
// or the other and reject the rest. The split is line-based and
// brace-balanced; the parser does the actual syntactic work.
func splitTopLevel(body string) (top, fnBody string) {
	lines := strings.Split(body, "\n")
	var topBuf, restBuf []string
	i, n := 0, len(lines)
	consumeBraced := func() {
		// collect lines until brace depth returns to 0 (or end-of-input).
		// Lines without any braces still count as one statement.
		depth := 0
		seenBrace := false
		for i < n {
			line := lines[i]
			topBuf = append(topBuf, line)
			depth += strings.Count(line, "{") - strings.Count(line, "}")
			if strings.ContainsAny(line, "{}") {
				seenBrace = true
			}
			i++
			if seenBrace && depth == 0 {
				return
			}
			if !seenBrace && strings.HasSuffix(strings.TrimRight(line, " \t"), ")") {
				// `var X = foo()` or single-line decl — stop on first `)`-terminated line.
				return
			}
		}
	}
	for i < n {
		line := lines[i]
		trim := strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(trim, "import ("):
			// multi-line import block — collect until matching `)`.
			topBuf = append(topBuf, line)
			i++
			for i < n {
				topBuf = append(topBuf, lines[i])
				done := strings.TrimSpace(lines[i]) == ")"
				i++
				if done {
					break
				}
			}
		case strings.HasPrefix(trim, "import "):
			topBuf = append(topBuf, line)
			i++
		case strings.HasPrefix(trim, "func "), strings.HasPrefix(trim, "var "),
			strings.HasPrefix(trim, "const "), strings.HasPrefix(trim, "type "),
			strings.HasPrefix(trim, "package "):
			consumeBraced()
		case trim == "":
			// blank line — attach to whichever buffer is currently
			// being filled. Once we've started filling restBuf, all
			// subsequent blank lines stay with rest.
			if len(restBuf) == 0 {
				topBuf = append(topBuf, line)
			} else {
				restBuf = append(restBuf, line)
			}
			i++
		default:
			restBuf = append(restBuf, line)
			i++
		}
	}
	return strings.Join(topBuf, "\n"), strings.Join(restBuf, "\n")
}

func main() {
	body := strings.TrimSpace(wrappeeBody)
	top, fn := splitTopLevel(body)
	src := "package wrappee\n\n" + top + "\n"
	if strings.TrimSpace(fn) != "" {
		src += "\nfunc _wrappee() {\n" + fn + "\n}\n"
	}
	fset := token.NewFileSet()
	if _, err := parser.ParseFile(fset, "fragment.go", src, parser.AllErrors); err != nil {
		fmt.Fprintln(os.Stderr, "parse error:", err)
		fmt.Fprintln(os.Stderr, "--- synthesized source ---")
		fmt.Fprintln(os.Stderr, src)
		os.Exit(1)
	}
	// EXAM-HELLO success line — emitted on a syntactically clean parse.
	fmt.Println("feature flag evaluates to true")
}
```
