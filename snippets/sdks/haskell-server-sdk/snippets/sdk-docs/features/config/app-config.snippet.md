---
id: haskell-server-sdk/sdk-docs/features/config/app-config
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Application metadata configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-config-syntax-only

---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.Config

import Data.Function ((&))

config :: Config
config = makeConfig "YOUR_SDK_KEY" & configSetApplicationInfo appInfo
    where appInfo = makeApplicationInfo
            & withApplicationValue "id" "authentication-service"
            & withApplicationValue "version" "1.0.0"
```
