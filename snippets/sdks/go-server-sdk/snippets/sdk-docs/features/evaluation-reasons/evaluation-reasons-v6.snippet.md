---
id: go-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Flag evaluation reason example for Go SDK v6+ (LDClient).
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
value, detail, err := client.BoolVariationDetail("example-flag-key", context, false)
// or StringVariationDetail for a string-valued flag, etc.

index := detail.VariationIndex
reason := detail.Reason
```
