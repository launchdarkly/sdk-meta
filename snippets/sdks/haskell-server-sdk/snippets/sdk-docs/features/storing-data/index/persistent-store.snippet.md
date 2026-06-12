---
id: haskell-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Persistent feature store configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-syntax-only-module

---

```haskell
import LaunchDarkly.Server

main = do
    backend <- makeYourBackendInterface

    let config = configSetStoreBackend backend $ makeConfig "YOUR_SDK_KEY"

    client <- makeClient config
```
