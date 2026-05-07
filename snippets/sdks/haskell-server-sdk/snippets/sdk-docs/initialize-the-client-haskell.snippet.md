---
id: haskell-server-sdk/sdk-docs/initialize-the-client-haskell
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: "Haskell in section \"Initialize the client\""
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only
---

```haskell
client :: IO Client
client = makeClient $ makeConfig "YOUR_SDK_KEY"
```
