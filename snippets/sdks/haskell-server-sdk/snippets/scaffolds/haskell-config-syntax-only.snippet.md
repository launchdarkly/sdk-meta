---
id: haskell-server-sdk/scaffolds/haskell-config-syntax-only
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  Parse-and-compile validator for Haskell config doc fragments that are
  already complete module-level Haskell — a leading `{-# LANGUAGE … #-}`
  pragma, `import` lines, and top-level bindings. Unlike haskell-syntax-only
  (which wraps an expression body in a `do` block and lifts top-level
  decls out via TOP_LIFT), these fragments ARE the module, so the scaffold
  just appends a `main` that prints the EXAM-HELLO line. The body's own
  pragma stays at file line 1 (Haskell requires LANGUAGE pragmas before any
  module/decl), its imports precede the appended main, and an implicit
  `Main` module is fine for a cabal executable. No TOP_LIFT marker, so the
  harness stages the body verbatim.
inputs:
  body:
    type: string
    description: A complete module-level Haskell fragment (pragma + imports + bindings).
validation:
  runtime: haskell-server
  entrypoint: app/Main.hs
---

```haskell
{{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
