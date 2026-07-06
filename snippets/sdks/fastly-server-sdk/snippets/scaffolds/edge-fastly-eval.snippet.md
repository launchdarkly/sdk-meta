---
id: fastly-server-sdk/scaffolds/edge-fastly-eval
sdk: fastly-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Fastly "evaluate a flag" fragment, which carries its own
  LDContext import and assumes an initialized `ldClient`. The
  _globals.d.ts companion ambiently types `ldClient` as the real return
  of the Fastly SDK's init, so variation(...) is checked against the real
  client API while the fragment's own import stays at module top.
inputs:
  body:
    type: string
    description: The wrappee fragment; uses the ambient `ldClient`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
  companions:
    - fastly-server-sdk/scaffolds/edge-fastly-eval-globals
---

```typescript
{{ body }}
```
