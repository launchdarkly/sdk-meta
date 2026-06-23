---
id: go-server-sdk/sdk-docs/features/privateattrs/context-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Marking context attributes private with the context builder for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
    "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
)

context := ldcontext.NewBuilder("example-context-key").
    Kind("organization").
    Name("Global Health Services").
    SetString("email", "info@globalhealthexample.com").
    SetValue("address", ldvalue.ObjectBuild().
        SetString("street", "123 Main Street").
        SetString("city", "Springfield").
        Build()).
    Private("email").
    Private("/address/street").
    Build()
```
