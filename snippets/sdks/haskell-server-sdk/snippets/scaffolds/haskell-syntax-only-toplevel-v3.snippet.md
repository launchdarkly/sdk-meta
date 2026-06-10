---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel-v3
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  v3.x flavor of `haskell-syntax-only-toplevel` — identical splice
  shape, but routes through the `haskell-server-v3` validator (which
  pins the 3.x package) and stubs `user` instead of `context`, since
  v3-era fragments pass a bare `user` to the variation functions.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at module scope.
validation:
  runtime: haskell-server-v3
  entrypoint: app/Main.hs
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
module Main where

import LaunchDarkly.Server

-- Module-scope stubs for the ambient bindings the doc fragments
-- assume earlier init snippets created.
client :: Client
client = undefined

user :: User
user = undefined

{{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
