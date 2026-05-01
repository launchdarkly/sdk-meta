---
id: go-server-sdk/sdk-docs/migration-5-to-6-understanding-changes-to-alias-events-6-0-syntax-associating-two-contexts
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```go
context1 := ldcontext.New("example-user-key")
context2 := ldcontext.NewWithKind("device", "example-device-key")
multiContext := ldcontext.NewMulti(context1, context2)
client.Identify(multiContext)
```
