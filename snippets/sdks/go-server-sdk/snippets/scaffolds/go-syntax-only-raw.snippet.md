---
id: go-server-sdk/scaffolds/go-syntax-only-raw
sdk: go-server-sdk
kind: scaffold
lang: go
file: fragment.txt
description: |
  Parse-only validator for Go server SDK doc fragments whose bodies
  contain backtick characters (for example Markdown-style code
  references inside comments). `go-syntax-only` embeds the wrappee
  body in a backtick-delimited Go raw string, so a body containing a
  backtick terminates the literal and breaks the scaffold itself.

  This variant stages the body verbatim as `fragment.txt` (this file)
  and ships the parser program as the `go-syntax-only-raw-main`
  companion, which reads the fragment from disk at run time. The
  split/parse semantics are identical to `go-syntax-only`: top-level
  declarations are spliced at file scope, statement residue is wrapped
  in a no-op function, and `go/parser.ParseFile` checks the result
  without resolving symbols.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, staged verbatim as fragment.txt.
validation:
  runtime: go
  entrypoint: main.go
  companions:
    - go-server-sdk/scaffolds/go-syntax-only-raw-main
---

```go
{{ body }}
```
