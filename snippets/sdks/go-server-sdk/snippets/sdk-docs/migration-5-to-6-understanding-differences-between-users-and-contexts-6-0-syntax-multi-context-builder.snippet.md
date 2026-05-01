---
id: go-server-sdk/sdk-docs/migration-5-to-6-understanding-differences-between-users-and-contexts-6-0-syntax-multi-context-builder
sdk: go-server-sdk
kind: reference
lang: go
description: "6.0 syntax, multi-context builder in section \"Understanding differences between users and contexts\""
---

```go
multiContext := ldcontext.NewMulti(
    ldcontext.NewBuilder("example-user-key").Name("Sandy").Build(),
    ldcontext.NewBuilder("example-device-key").Kind("device").Name("iPad").Build(),
)
```
