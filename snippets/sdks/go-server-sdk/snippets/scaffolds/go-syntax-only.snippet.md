---
id: go-server-sdk/scaffolds/go-syntax-only
sdk: go-server-sdk
kind: scaffold
lang: go
file: main.go
description: |
  Parse-only validator for Go server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: go
  entrypoint: main.go
---

```go
package main

import (
	"fmt"

	_ "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
	_ "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
	_ "github.com/launchdarkly/go-server-sdk/v7"
)

func main() {
	fmt.Println("feature flag evaluates to true")
}

// _wrappee compiles the doc-fragment body in a function the harness never
// invokes — the EXAM-HELLO success line is printed unconditionally above,
// so we just need the snippet to parse and link.
func _wrappee() {
{{ body }}
}
```
