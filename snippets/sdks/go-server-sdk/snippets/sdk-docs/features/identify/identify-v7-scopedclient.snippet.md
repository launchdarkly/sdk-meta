---
id: go-server-sdk/sdk-docs/features/identify/identify-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Identify example for the Go SDK v7.13.4+, using LDScopedClient.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
// There is not an Identify method in the LDScopedClient,
// so you need to access the method from the LDClient.
// Then, pass in the scoped client's current context.
// LDScopedClient is in beta and may change without notice.
scopedClient.Client().Identify(scopedClient.CurrentContext())
```
