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
-- Qualified helpers several doc fragments use without showing their
-- own import lines (the docs assume an earlier snippet added them).
import qualified Data.Set as S
import qualified LaunchDarkly.Server.Reference as R
import Data.Function ((&))
import Data.Text (Text)

-- The body is spliced at module scope, BEFORE the stub declarations
-- below, so fragments that carry their own leading `import` lines
-- stay legal (Haskell requires every import to precede the module's
-- first declaration; declaration order itself is irrelevant).
{{ body }}

-- Module-scope stubs for the ambient bindings the doc fragments
-- assume earlier init snippets created.
client :: Client
client = undefined

context :: Context
context = undefined

-- Config fragments pass a bare `sdkKey` the docs assume exists.
sdkKey :: Text
sdkKey = "YOUR_SDK_KEY"

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
