---
id: haskell-server-sdk/scaffolds/haskell-syntax-only-module
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Compile-only validator for Haskell doc fragments that are a complete
  module body INCLUDING a `main = do` binding (the storing-data
  fragments, for example). `haskell-config-syntax-only` cannot host
  these (it appends its own `main`, which would collide), and running
  the body's `main` would try to reach a real database, so the
  scaffold sets `SNIPPET_CHECK=parse` and the harness stops after a
  clean `cabal build`.

  Two quirks of the splice:

  - The fragments end their `main` do-block with a bind statement
    (`client <- makeClient config`) — a doc idiom inviting the reader
    to continue, but GHC requires a do-block to end with an
    expression. The scaffold appends a `pure ()` line at the
    fragments' 4-space statement indentation to complete the block,
    keeping the published body verbatim.
  - `makeYourBackendInterface` is the docs' generic placeholder for a
    store-integration constructor. The stub is intentionally left
    without a type signature: `undefined` generalizes, so the same
    stub serves whatever store type the body's use site infers.
inputs:
  body:
    type: string
    description: A complete module-level Haskell fragment (imports + bindings, defining main).
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
  env:
    SNIPPET_CHECK: parse
---

```haskell
{-# LANGUAGE OverloadedStrings #-}
{{ body }}
    pure ()

-- Generic database-integration placeholder the storing-data docs use.
makeYourBackendInterface = undefined
```
