---
id: haskell-server-sdk/sdk-docs/features/contextconfig/multi-context-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Multi-context example for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-expression

---

```haskell
makeMultiContext [ makeContext "example-user-key" "user"
                 , makeContext "example-device-key" "device"
                 ]
```
