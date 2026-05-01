---
id: haskell-server-sdk/sdk-docs/migration-3-to-4-working-with-built-in-and-custom-attributes-4-0-syntax-context-with-attributes
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "4.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```haskell
makeContext "example-user-key" "user"
  & withName "Sandy"
  & withAttribute "email" "sandy@example.com"
```
