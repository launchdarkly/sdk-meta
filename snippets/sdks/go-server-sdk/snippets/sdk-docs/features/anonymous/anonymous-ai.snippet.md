---
id: go-server-sdk/sdk-docs/features/anonymous/anonymous-ai
sdk: go-server-sdk
kind: reference
lang: go
description: Anonymous context example for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
)

// Anonymous context with only a key
context1 := ldcontext.NewBuilder("example-context-key").Anonymous(true)

// Anonymous context with a key plus other attributes
context2 := ldcontext.NewBuilder("context-key-456def").
    Anonymous(true).
    SetString("country", "Canada").
    Build()
```
