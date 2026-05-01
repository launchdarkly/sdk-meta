---
id: go-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-private-attributes-6-0-syntax-attribute-marked-private-for-one-context
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```go
context := ldcontext.NewBuilder("key").
    Name("Sandy").
    SetString("email", "sandy@example.com").
    Private("email").
    Build()
```
