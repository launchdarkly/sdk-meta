---
id: go-server-sdk/sdk-docs/migration-5-to-6-working-with-built-in-and-custom-attributes-6-0-syntax-context-with-attributes
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```go
context2 := ldcontext.NewBuilder("example-user-key").
    Name("Sandy").
    SetString("email", "sandy@example.com").
    SetValue("groups", ldvalue.ArrayOf(
      ldvalue.String("Acme"), ldvalue.String("Global Health Services"))).
    Build()
```
