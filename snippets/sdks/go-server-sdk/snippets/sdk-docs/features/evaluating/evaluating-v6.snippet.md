---
id: go-server-sdk/sdk-docs/features/evaluating/evaluating-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Flag evaluation example for Go.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only

---

```go
result, _ := client.BoolVariation("example-flag-key", context, false)

// result is now true or false depending on the setting of this boolean feature flag
```
