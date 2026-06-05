---
id: haskell-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
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
    & configSetStreamURI "https://your-relay-proxy.com:8030"
    & configSetBaseURI "https://your-relay-proxy.com:8030"
    & configSetEventsURI "https://your-relay-proxy.com:8030"
```
