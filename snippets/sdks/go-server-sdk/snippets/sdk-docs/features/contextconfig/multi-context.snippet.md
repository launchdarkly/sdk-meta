---
id: go-server-sdk/sdk-docs/features/contextconfig/multi-context
sdk: go-server-sdk
kind: reference
lang: go
description: Multi-context example for Go SDK v6+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
multiContext := ldcontext.NewMulti(
    ldcontext.New("example-user-key"),
    ldcontext.NewWithKind("device", "example-device-key"),
)
```
