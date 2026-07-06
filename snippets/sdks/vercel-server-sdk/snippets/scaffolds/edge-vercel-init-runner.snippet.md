---
id: vercel-server-sdk/scaffolds/edge-vercel-init-runner
sdk: vercel-server-sdk
kind: scaffold
lang: typescript
file: snippet.ts
description: |
  Type-checks a Vercel "initialize the client" fragment. Supplies the
  `init` and `createClient` imports the page's import block documents,
  then splices the fragment, which builds the edge-config client and
  calls `init(...)`. tsc checks both calls against the real signatures.
inputs:
  body:
    type: string
    description: The wrappee init fragment; binds edgeConfigClient and ldClient.
validation:
  runtime: edge-tsc
  entrypoint: snippet.ts
---

```typescript
import { init } from '@launchdarkly/vercel-server-sdk';
import { createClient } from '@vercel/edge-config';

{{ body }}
```
