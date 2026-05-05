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
  module level. The harness pre-stage rewrite (in run.sh) splits the
  body into top-level lines vs in-block lines using a simple
  indentation rule: any line that begins with `import `, `data `,
  `type `, `newtype `, `class `, `instance `, or matches `name :: …`
  (a type-signature line) is lifted to the TOP_LIFT_MARKER above the
  module, along with any subsequent line that starts at column 0
  (top-level binding for that name). Everything else stays in the
  `do` block.
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

import qualified Data.Function as _LD

-- TOP_LIFT_MARKER

main :: IO ()
main = putStrLn "feature flag evaluates to true"

_wrappee :: IO ()
_wrappee = do
{{ body }}
  return ()
```
