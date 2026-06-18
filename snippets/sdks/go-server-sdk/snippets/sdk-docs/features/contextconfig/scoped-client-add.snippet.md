---
id: go-server-sdk/sdk-docs/features/contextconfig/scoped-client-add
sdk: go-server-sdk
kind: reference
lang: go
description: Adding contexts to a scoped client for Go SDK v7.13.4+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
userContext := ldcontext.New("example-user-key")

scopedClient := ld.NewScopedClient(client, userContext)
scopedClient.CurrentContext() // returns the single "user" context

scopedClient.AddContext(ldcontext.NewWithKind("device", "example-device-key"))
scopedClient.CurrentContext() // returns a multi-context with "user" and "device" contexts

scopedClient.BoolVariation("example-flag-key", false) // evaluates the flag using a multi-context with "user" and "device" contexts
```
