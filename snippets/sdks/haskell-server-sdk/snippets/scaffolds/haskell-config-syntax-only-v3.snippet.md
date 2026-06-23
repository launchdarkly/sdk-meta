---
id: haskell-server-sdk/scaffolds/haskell-config-syntax-only-v3
sdk: haskell-server-sdk
kind: scaffold
lang: haskell
file: app/Main.hs
description: |
  v3.x flavor of `haskell-config-syntax-only` — same splice shape
  (the fragment IS the module: leading `{-# LANGUAGE ... #-}` pragma,
  `import` lines, top-level bindings; the scaffold just appends a
  `main`), but routes through the `haskell-server-v3` validator so
  fragments referencing the 3.x `LaunchDarkly.Server.User` surface
  compile against the actual v3.x package.
inputs:
  body:
    type: string
    description: A complete module-level Haskell fragment (pragma + imports + bindings).
validation:
  runtime: haskell-server-v3
  entrypoint: app/Main.hs
---

```haskell
{{ body }}

main :: IO ()
main = putStrLn "feature flag evaluates to true"
```
