---
id: haskell-server-sdk/sdk-docs/features/contextconfig/context-example-v3
sdk: haskell-server-sdk
kind: reference
lang: haskell
description: User example for Haskell SDK v3.x.
validation:
  scaffold: haskell-server-sdk/scaffolds/haskell-config-syntax-only-v3

---

```haskell
{-# LANGUAGE OverloadedStrings #-}

import LaunchDarkly.Server.User

import Data.Function ((&))

-- User with only a key
user1 :: User
user1 = (makeUser "example-user-key")

user2 :: User
user2 = (makeUser "user-key-456def")
    & userSetFirstName (Just "Sandy")
    & userSetLastName (Just "Smith")
    & userSetEmail (Just "sandy@example.com")
```
