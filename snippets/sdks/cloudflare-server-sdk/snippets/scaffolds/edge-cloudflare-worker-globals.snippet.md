---
id: cloudflare-server-sdk/scaffolds/edge-cloudflare-worker-globals
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: _globals.d.ts
description: |
  Ambient companion for edge-cloudflare-worker. Declares the `Bindings`
  type the Worker's fetch handler annotates `env` with, so the fragment
  stays verbatim (its own import at module top) while `env.LD_KV`
  type-checks as a real KVNamespace.
---

```typescript
type Bindings = { LD_KV: Parameters<typeof import('@launchdarkly/cloudflare-server-sdk').init>[1] };
```
