---
id: haskell-server-sdk/sdk-docs/features/testdata/set-flag-value-v3x
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Setting a test data flag to a specific value for Haskell SDK v3.x.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-v3
---

```haskell
TestData.update td =<< ( TestData.flag td "example-flag-key"
    <&> TestData.booleanFlag
    <&> TestData.variationForAllUsers True
)
```
