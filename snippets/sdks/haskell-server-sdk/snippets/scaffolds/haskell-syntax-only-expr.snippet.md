---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-expr
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Compile validator for Haskell doc fragments that are a bare pure
  expression (e.g. `secureModeHash client context` or
  `makeContext "key" "user" & withAnonymous True`).

  `haskell-syntax-only` places the body inside `_wrappee`'s do-block,
  where a non-IO expression statement is a type error, and the
  toplevel variants splice the body at module scope, where a bare
  expression is not a valid declaration. This variant binds the
  expression to a module-level `_wrappee` name instead. The splice
  sits on the binding's right-hand side, so continuation lines that
  the docs indent (e.g. a leading `& withAnonymous True`) stay
  layout-valid. Module-scope `client` / `context` stubs mirror
  `haskell-syntax-only-toplevel`, and `Data.Function ((&))` is
  imported because the doc fragments use `&` chaining without showing
  the import.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, a pure expression bound to a module-level name.
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
module Main where

import LaunchDarkly.Server
import Data.Function ((&))

-- Module-scope stubs for the ambient bindings the doc fragments
-- assume earlier init snippets created.
client :: Client
client = undefined

context :: Context
context = undefined

_wrappee = {{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
