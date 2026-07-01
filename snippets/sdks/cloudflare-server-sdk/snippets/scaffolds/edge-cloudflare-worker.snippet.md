---
id: cloudflare-server-sdk/scaffolds/edge-cloudflare-worker
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks the Cloudflare example Worker fragment, a whole module with
  its own imports and an exported fetch handler. The _globals.d.ts
  companion declares the `Bindings` type the handler annotates its `env`
  with (LD_KV as a real KVNamespace), so the whole handler type-checks
  against the real SDK and Workers types.
inputs:
  body:
    type: string
    description: The wrappee Worker module, spliced verbatim.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
  companions:
    - cloudflare-server-sdk/scaffolds/edge-cloudflare-worker-globals
---

```typescript
{{ body }}
```
