---
id: go-server-sdk/scaffolds/go-syntax-only-raw-main
sdk: go-server-sdk
kind: scaffold
lang: go
file: main.go
description: |
  Parser half of the `go-syntax-only-raw` scaffold pair. Staged as a
  companion so it lands at `main.go` while the wrappee body sits
  verbatim at `fragment.txt`. Reading the fragment from disk (instead
  of embedding it in a raw string the way `go-syntax-only` does) keeps
  bodies containing backticks parseable. The split logic mirrors
  `go-syntax-only`; keep the two in sync.
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
		// Read until a top-level declaration is complete.
		// A decl ends when whichever delimiter style it uses closes:
		//   - `var foo struct { ... }`  — brace-delimited; ends when
		//     brace depth returns to 0 after at least one brace seen.
		//   - `var ( foo Type\n )`      — paren-delimited; ends when
		//     paren depth returns to 0 after at least one paren seen.
		//   - `var foo SomeType`        — neither braces nor parens; the
		//     decl is the single first line.
		// Tracking braces alone leaks subsequent statements into the
		// top-level bucket when the decl has no braces but the body
		// continues with statements (e.g. `var x T` followed by
		// `x.field = …`).
		bdepth, pdepth := 0, 0
		seenBrace, seenParen := false, false
		for i < n {
			line := lines[i]
			topBuf = append(topBuf, line)
			bdepth += strings.Count(line, "{") - strings.Count(line, "}")
			pdepth += strings.Count(line, "(") - strings.Count(line, ")")
			if strings.ContainsAny(line, "{}") {
				seenBrace = true
			}
			if strings.ContainsAny(line, "()") {
				seenParen = true
			}
			i++
			if seenBrace && bdepth == 0 {
				return
			}
			if !seenBrace && seenParen && pdepth == 0 {
				return
			}
			if !seenBrace && !seenParen {
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
	raw, err := os.ReadFile("fragment.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "scaffold: cannot read fragment.txt:", err)
		os.Exit(1)
	}
	body := strings.TrimSpace(string(raw))
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
