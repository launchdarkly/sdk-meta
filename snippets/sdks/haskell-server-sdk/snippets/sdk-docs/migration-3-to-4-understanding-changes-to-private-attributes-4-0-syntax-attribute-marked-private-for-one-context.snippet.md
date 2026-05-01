---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-understanding-changes-to-private-attributes-4-0-syntax-attribute-marked-private-for-one-context
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "4.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```haskell
makeContext "key" "user"
  & withName "Sandy"
  & withAttribute "email" "sandy@example.com"
  & withPrivateAttributes (S.fromList $ map R.makeLiteral ["name", "email"])
```
