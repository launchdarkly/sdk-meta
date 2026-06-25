---
id: go-server-sdk/sdk-docs/features/testdata/set-flag-value-v6
sdk: go-server-sdk
kind: reference
lang: go
description: Setting a test data flag to a specific value for Go SDK v6.0.
validation:
  scaffold: go-server-sdk/scaffolds/go-syntax-only
---

```go
td.Update(td.Flag("example-flag-key").VariationForAll(false))
```
