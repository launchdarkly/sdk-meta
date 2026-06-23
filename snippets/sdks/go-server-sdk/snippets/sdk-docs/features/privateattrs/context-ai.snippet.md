---
id: go-server-sdk/sdk-docs/features/privateattrs/context-ai
sdk: go-server-sdk
kind: reference
lang: go
description: Marking context attributes private with the context builder for the Go AI SDK.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
)

context := ldcontext.NewBuilder("example-context-key").
    Kind("organization").
    Name("Global Health Services").
    SetString("email", "info@globalhealthexample.com").
    Private("email").
    Build()
```
