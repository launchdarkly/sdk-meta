---
id: haskell-server-sdk/scaffolds/haskell-syntax-only
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Parse-only validator for Haskell server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
---

```haskell
module Main where

main :: IO ()
main = putStrLn "feature flag evaluates to true"

_wrappee :: IO ()
_wrappee = do
{{ body }}
```
