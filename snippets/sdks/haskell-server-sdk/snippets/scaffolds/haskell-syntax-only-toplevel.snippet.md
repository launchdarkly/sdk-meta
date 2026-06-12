---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-toplevel
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Module-scope sibling of `haskell-syntax-only`. That scaffold places
  the body inside `_wrappee`'s do-block (with let-bound `client` /
  `context` stubs), and the harness lifts top-level-shaped lines out
  to module scope — where the do-block stubs aren't visible. Bodies
  that are entirely top-level declarations referencing an ambient
  `client` / `context` (e.g. `details :: IO (EvaluationDetail Bool)`
  with its binding) therefore need the stubs at module scope, which
  is what this variant provides. The body is spliced directly at
  module scope, so no lift markers are involved.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, spliced at module scope.
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

-- Config fragments derive a new config from an ambient `config`
-- binding the docs assume an earlier init snippet created.
config :: Config
config = undefined

{{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
