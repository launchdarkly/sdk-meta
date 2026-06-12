---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-expression
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Sibling of `haskell-syntax-only` for doc fragments that are a bare
  `Context` expression (e.g. `makeContext "key" "organization"` or a
  `makeMultiContext [...]` call). A bare non-IO expression is neither
  valid at module scope nor as a statement in the do-block scaffold,
  so this variant binds the expression to a module-scope `_fragment`
  name and lets GHC type-check it from there. Multi-line bodies work
  because their continuation lines are already indented past column 1.
inputs:
  body:
    type: string
    description: A Haskell expression of type Context.
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
module Main where

import LaunchDarkly.Server

-- The doc fragment is a bare Context expression; binding it at
-- module scope lets GHC type-check it without a surrounding do-block.
_fragment :: Context
_fragment = {{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
