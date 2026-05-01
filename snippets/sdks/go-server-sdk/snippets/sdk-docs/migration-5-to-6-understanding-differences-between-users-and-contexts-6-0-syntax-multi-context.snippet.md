---
id: go-server-sdk/sdk-docs/migration-5-to-6-understanding-differences-between-users-and-contexts-6-0-syntax-multi-context
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```go
multiContext := ldcontext.NewMulti(
    ldcontext.New("example-user-key"),
    ldcontext.NewWithKind("device", "example-device-key")
)
```
