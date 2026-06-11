---
id: haskell-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Flag evaluation reason example for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel

---

```haskell
details :: IO (EvaluationDetail Bool)
details = boolVariationDetail client "example-flag-key" context False
```
