---
id: haskell-server-sdk/sdk-docs/features/config/index
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: SDK configuration example for Haskell.
---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.Config

import Data.Function ((&))

config :: Config
config = (makeConfig "YOUR_SDK_KEY")
    & configSetEventsCapacity 1000
    & configSetFlushIntervalSeconds 30
```
