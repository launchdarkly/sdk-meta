---
id: haskell-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Service endpoint configuration example for Haskell.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-config-syntax-only

---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.Config

import Data.Function ((&))

config :: Config
config = (makeConfig "YOUR_SDK_KEY")
    & configSetStreamURI "https://stream.launchdarkly.us"
    & configSetBaseURI "https://sdk.launchdarkly.us"
    & configSetEventsURI "https://events.launchdarkly.us"
```
