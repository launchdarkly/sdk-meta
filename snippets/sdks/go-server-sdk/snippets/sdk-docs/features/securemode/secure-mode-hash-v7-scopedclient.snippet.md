---
id: go-server-sdk/sdk-docs/features/securemode/secure-mode-hash-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Secure mode hash example for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
// There is not a SecureModeHash method in the LDScopedClient,
// so you need to access the method from the LDClient.
// Then, pass in the scoped client's current context.
// LDScopedClient is in beta and may change without notice.
scopedClient.Client().SecureModeHash(scopedClient.CurrentContext())
```
