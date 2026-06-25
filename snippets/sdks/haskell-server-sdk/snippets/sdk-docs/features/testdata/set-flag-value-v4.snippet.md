---
id: haskell-server-sdk/sdk-docs/features/testdata/set-flag-value-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Setting a test data flag to a specific value for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only
---

```haskell
TestData.update td =<< ( TestData.flag td "example-flag-key"
	<&> TestData.booleanFlag
	<&> TestData.variationForAll True
)
```
