---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-v3
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Parse-and-type-check validator for Haskell server SDK doc fragments
  that target the v3.x API surface (e.g. `makeUser`, `User`).

  Routes through the `haskell-server-v3` validator container, which
  pins `launchdarkly-server-sdk == 3.1.1` in its cabal file. The
  current-version `haskell-syntax-only` scaffold compiles against
  the latest SDK and won't resolve removed names like `makeUser`;
  this scaffold exists so v3.x-specific docs validate against the
  actual v3.x SDK that the docs cover.

  The wrappee body is dropped inside `_wrappee = do { … }`. Same
  harness pre-stage rewrite as `haskell-syntax-only`: top-level
  declarations (imports, data, type, newtype, class, instance,
  type sigs, top-level bindings) get lifted to the
  `--TOP_LIFT_TARGET--` marker at module scope; everything else
  stays inside the do-block.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: haskell-server-v3
  entrypoint: app/Main.hs
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
module Main where

import LaunchDarkly.Server
import qualified Data.Function as LDStub

-- File-scope stubs so wrappee bodies that reference `client` (the
-- SDK client value the consumer constructs in their own application
-- code) resolve at parse time. The wrappee body never executes;
-- `undefined` is sufficient for the type-checker.
client :: Client
client = undefined

--TOP_LIFT_TARGET--

main :: IO ()
main = putStrLn "feature flag evaluates to true"

-- Body lives inside a do-block. The harness BB / BE markers
-- delimit body content; the awk pre-step lifts top-level decls
-- found there to the target marker above.
_wrappee :: IO ()
_wrappee = do
--BODY_BEGIN--
{{ body }}
--BODY_END--
  return ()
```
