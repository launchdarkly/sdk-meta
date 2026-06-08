---
id: go-server-sdk/sdk-docs/features/evaluating/evaluating-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Flag evaluation example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
result, _ := scopedClient.BoolVariation("example-flag-key", false)

// result is now true or false depending on the setting of this boolean feature flag
// LDScopedClient is in beta and may change without notice.
```
