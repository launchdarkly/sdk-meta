---
id: haskell-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-v3
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Flag evaluation reason example for Haskell SDK v3.x.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel-v3

---

```haskell
details :: IO (EvaluationDetail Bool)
details = boolVariationDetail client "example-flag-key" user False
```
