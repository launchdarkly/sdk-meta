---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-working-with-built-in-and-custom-attributes-3-x-syntax-user-with-attributes
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "3.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```haskell
makeUser "example-user-key"
 & userSetName (Just "Sandy")
 & userSetEmail (Just "sandy@example.com")
```
