---
id: haskell-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Daemon mode configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-config-syntax-only
---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.Config

import Data.Function ((&))

config :: Config
config = (makeConfig "YOUR_SDK_KEY")
    & configSetUseLdd True
    & configSetStoreBackend backend
```
