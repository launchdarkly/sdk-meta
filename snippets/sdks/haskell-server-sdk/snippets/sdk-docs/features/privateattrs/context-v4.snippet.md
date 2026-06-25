---
id: haskell-server-sdk/sdk-docs/features/privateattrs/context-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Marking context attributes private for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel

---

```haskell
context' = makeContext "key" "user"
  & withName "Sandy"
  & withAttribute "email" "sandy@example.com"
  & withPrivateAttributes (S.fromList $ map R.makeLiteral ["name", "email"])
```
