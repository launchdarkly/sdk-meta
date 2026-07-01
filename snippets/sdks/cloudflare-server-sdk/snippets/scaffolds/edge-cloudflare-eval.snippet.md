---
id: cloudflare-server-sdk/scaffolds/edge-cloudflare-eval
sdk: cloudflare-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Cloudflare "evaluate a flag" fragment, which assumes an
  initialized `client`. The scaffold supplies a real `client`, then splices
  the fragment so `client.variation(...)` is checked against the real API.
inputs:
  body:
    type: string
    description: The wrappee fragment; assumes `client`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/cloudflare-server-sdk';

declare const env: { LD_KV: Parameters<typeof init>[1] };
const client = init('example-client-side-id', env.LD_KV);

{{ body }}
```
