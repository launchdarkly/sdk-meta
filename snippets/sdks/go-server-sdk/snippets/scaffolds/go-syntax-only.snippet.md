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
	_ ld "github.com/launchdarkly/go-server-sdk/v7"
)

func _wrappee() {
{{ body }}
}

func main() {
	fmt.Println("feature flag evaluates to true")
}
```
