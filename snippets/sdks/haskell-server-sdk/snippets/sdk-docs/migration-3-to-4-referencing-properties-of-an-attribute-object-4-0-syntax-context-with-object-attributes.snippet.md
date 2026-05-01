---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-referencing-properties-of-an-attribute-object-4-0-syntax-context-with-object-attributes
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "4.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```haskell
makeContext "example-user-key" "user"
  & withName "Sandy"
  & withAttribute "address" (Object $ fromList [("street", "Main St"), ("city", "Springfield")])
```
