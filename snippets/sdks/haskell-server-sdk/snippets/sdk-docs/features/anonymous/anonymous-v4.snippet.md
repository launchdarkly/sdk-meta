---
id: haskell-server-sdk/sdk-docs/features/anonymous/anonymous-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Anonymous context example for Haskell, SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-expr
---

```haskell
makeContext "example-user-key" "user"
  & withAnonymous True
```
