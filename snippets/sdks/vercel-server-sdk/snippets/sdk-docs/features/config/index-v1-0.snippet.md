---
id: vercel-server-sdk/sdk-docs/features/config/index-v1-0
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Vercel.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
import { createClient } from '@vercel/edge-config';
import { BasicLogger, init, LDOptions } from '@launchdarkly/vercel-server-sdk';

const edgeConfigClient = createClient(process.env.EDGE_CONFIG);

const options: LDOptions = {
  logger: new BasicLogger({ destination: console.log }),
};

client = init('example-client-side-id', edgeConfigClient, options);
```
