---
id: haskell-server-sdk/sdk-docs/features/contextconfig/context-example-v4
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: Context example for Haskell SDK v4.0.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-config-syntax-only

---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.Context

import Data.Aeson (Value (Object))
import Data.Aeson.KeyMap (fromList)
import Data.Function ((&))

-- Context with key and kind
context1 :: Context
context1 = makeContext "example-context-key" "user"

-- Context with a key plus other attributes
context2 :: Context
context2 = makeContext "context-key-456def" "organization"
    & withAttribute "name" "Global Health Services"
    & withAttribute "email" "info@globalhealthexample.com"
    & withAttribute "address" (Object $ fromList [("street", "123 Main St"), ("city", "Springfield")])
```
