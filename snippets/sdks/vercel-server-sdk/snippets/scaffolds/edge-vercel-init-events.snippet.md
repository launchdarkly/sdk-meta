---
id: vercel-server-sdk/scaffolds/edge-vercel-init-events
sdk: vercel-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks the Vercel init-with-events fragment, which assumes an
  `edgeConfigClient` created in the previous doc block. The scaffold
  supplies the imports and a real `edgeConfigClient`, then splices the
  fragment so `init(..., { sendEvents: true })` is checked against the
  real overload.
inputs:
  body:
    type: string
    description: The wrappee fragment; assumes `edgeConfigClient` and binds `ldClient`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/vercel-server-sdk';
import { createClient } from '@vercel/edge-config';

const edgeConfigClient = createClient(process.env.EDGE_CONFIG);

{{ body }}
```
