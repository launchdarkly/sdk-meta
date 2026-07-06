---
id: cloudflare-server-sdk/scaffolds/edge-cloudflare-init-runner
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Cloudflare "initialize the client" fragment. Supplies the
  documented `init` import and a Workers `env` binding whose `LD_KV` is a
  real `KVNamespace` (from @cloudflare/workers-types), so `init(clientId,
  env.LD_KV)` is checked against the real signature.
inputs:
  body:
    type: string
    description: The wrappee init fragment; calls init(...) with env.LD_KV and binds client.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/cloudflare-server-sdk';

declare const env: { LD_KV: Parameters<typeof init>[1] };

{{ body }}
```
