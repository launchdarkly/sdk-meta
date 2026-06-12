---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-expr
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Expression-scope sibling of `haskell-syntax-only`. That scaffold
  places the body inside `_wrappee`'s do-block, where each body line
  must be a monadic statement — but some doc fragments are a bare pure
  expression (e.g. `secureModeHash client context :: Text`), which
  GHC rejects as a do-statement. This variant splices the body as the
  right-hand side of a module-scope binding instead, so the fragment
  type-checks as the expression the docs present it as. Module-scope
  `client` / `context` stubs mirror `haskell-syntax-only-toplevel`.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced as the right-hand side of a module-scope binding.
validation:
  runtime: haskell-server
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

context :: Context
context = undefined

_wrappeeExpr = {{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
