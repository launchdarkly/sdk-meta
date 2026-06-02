---
id: haskell-server-sdk/scaffolds/haskell-syntax-only
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Parse-only validator for Haskell server SDK doc fragments.

  The body is dropped inside `_wrappee = do { … }`. Doc fragments that
  show top-level constructs (`import …`, `name :: Type`, top-level
  bindings) won't fit inside a `do` block — Haskell wants those at the
  module level. The harness pre-stage rewrite (in `run.sh`) splits the
  body using `--BODY_BEGIN--` / `--BODY_END--` markers: any body line
  at column 0 that begins with a top-level keyword (`import `, `data `,
  `type `, `newtype `, `class `, `instance `) or matches a type-sig /
  binding shape (`name ::` or `name =`) is lifted to the
  `--TOP_LIFT_TARGET--` marker at module scope. Body residue stays
  inline but is indented two spaces so it lands inside the do-block.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
module Main where

import LaunchDarkly.Server
import qualified Data.Function as LDStub

--TOP_LIFT_TARGET--

main :: IO ()
main = putStrLn "feature flag evaluates to true"

-- Body lives inside a do-block. The harness BB / BE markers
-- delimit body content; the awk pre-step lifts top-level decls
-- found there to the target marker above.
_wrappee :: IO ()
_wrappee = do
  -- Local stub for bodies that reference bare `client` without
  -- declaring it themselves (e.g. evaluate-a-context fragments).
  -- Scoped to the do-block so it does NOT collide with bodies that
  -- declare their own top-level `client :: IO Client` (e.g.
  -- initialize-the-client). Such top-level declarations land via
  -- TOP_LIFT_TARGET outside the do-block.
  let client = undefined :: Client
--BODY_BEGIN--
{{ body }}
--BODY_END--
  return ()
```
