---
id: go-server-sdk/sdk-docs/features/overrides/override-affected
sdk: go-server-sdk
kind: reference
lang: go
description: Check whether a flag override affected an evaluation in Go SDK v7.18+.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
_, detail, _ := client.BoolVariationDetail("my-boolean-flag-key", context, false)
if detail.Reason.IsOverrideAffected() {
    fmt.Println("An override affected this evaluation")
}
```
