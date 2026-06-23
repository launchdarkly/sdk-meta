---
id: go-server-sdk/sdk-docs/features/contextconfig/multi-context-builder
sdk: go-server-sdk
kind: reference
lang: go
description: Multi-context example using the builder for Go SDK v6+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
multiContext := ldcontext.NewMulti(
    ldcontext.NewBuilder("example-user-key").Name("Sandy").Build(),
    ldcontext.NewBuilder("example-device-key").Kind("device").Name("iPad").Build(),
)
```
