---
id: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only
sdk: akamai-server-edgekv-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Parse-only validator for Akamai edge SDK config doc fragments. Routes
  through the `edge-ts` validator, which runs the TypeScript compiler's
  transpileModule (syntax check + type-strip; no module resolution, no
  type-checking). A clean parse means the fragment is syntactically
  valid TypeScript -- edge-only package imports and ambient globals
  (env, process) need not resolve. Doc fragments are whole TS modules,
  so the scaffold emits the body verbatim.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, parsed as a TypeScript module.
validation:
  runtime: edge-ts
  entrypoint: snippet.ts
---

```typescript
{{ body }}
```
