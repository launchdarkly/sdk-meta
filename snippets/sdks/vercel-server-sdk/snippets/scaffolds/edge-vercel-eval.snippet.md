---
id: vercel-server-sdk/scaffolds/edge-vercel-eval
sdk: vercel-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Vercel "evaluate a flag" fragment, which assumes an
  initialized `ldClient`. The scaffold supplies a real `ldClient`, then
  splices the fragment so `ldClient.variation(...)` is checked against
  the real client API.
inputs:
  body:
    type: string
    description: The wrappee fragment; assumes `ldClient`.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/vercel-server-sdk';
import { createClient } from '@vercel/edge-config';

const ldClient = init('example-client-side-id', createClient(process.env.EDGE_CONFIG));

{{ body }}
```
