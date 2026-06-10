---
id: go-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v7-scopedclient
sdk: go-server-sdk
kind: reference
lang: go
description: Flag evaluation reason example for Go SDK v7.13.4+ (LDScopedClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
value, detail, err := scopedClient.BoolVariationDetail("example-flag-key", false)
// or StringVariationDetail for a string-valued flag, etc.
// LDScopedClient is in beta and may change without notice.

index := detail.VariationIndex
reason := detail.Reason
```
