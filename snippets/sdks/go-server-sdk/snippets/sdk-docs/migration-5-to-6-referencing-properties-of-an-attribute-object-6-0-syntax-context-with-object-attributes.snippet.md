---
id: go-server-sdk/sdk-docs/migration-5-to-6-referencing-properties-of-an-attribute-object-6-0-syntax-context-with-object-attributes
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```go
addressData := ldvalue.ObjectBuild().
    SetString("street", "Main St").
    SetString("city", "Springfield").
    Build()
context3 := ldcontext.NewBuilder("example-user-key").
    SetValue("address", addressData).
    Build()
```
