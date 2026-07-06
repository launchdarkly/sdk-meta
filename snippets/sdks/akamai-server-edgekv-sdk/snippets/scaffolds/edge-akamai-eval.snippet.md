---
id: akamai-server-edgekv-sdk/scaffolds/edge-akamai-eval
sdk: akamai-server-edgekv-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks an Akamai "evaluate a flag" fragment against the real
  package. The fragment carries its own imports and calls
  `ldClient.variation(...)`; the `_globals.d.ts` companion ambiently
  declares `ldClient` as the real return type of `init`, so
  `variation(...)` is checked against the real client API. The fragment
  is spliced verbatim (its own import stays at module top).
inputs:
  body:
    type: string
    description: The wrappee fragment; uses the ambient `ldClient`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
  companions:
    - akamai-server-edgekv-sdk/scaffolds/edge-akamai-eval-globals
---

```typescript
{{ body }}
```
