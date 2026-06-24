---
id: go-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: go-server-sdk
kind: reference
lang: go
description: Context example for Go SDK v6+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
import (
    "github.com/launchdarkly/go-sdk-common/v3/ldcontext"
    "github.com/launchdarkly/go-sdk-common/v3/ldvalue"
)

// Context with only a key
// by default, the context kind is "user"
context1 := ldcontext.New("example-context-key")

// Context with a key plus other attributes
context2 := ldcontext.NewBuilder("context-key-456def").
    Kind("organization").
    Name("Global Health Services").
    SetString("email", "info@globalhealthexample.com").
    SetValue("address", ldvalue.ObjectBuild().
        SetString("street", "123 Main Street").
        SetString("city", "Springfield")).
    SetValue("groups", ldvalue.ArrayOf(
      ldvalue.String("Acme"), ldvalue.String("Global Health Services"))).
    Build()
```
