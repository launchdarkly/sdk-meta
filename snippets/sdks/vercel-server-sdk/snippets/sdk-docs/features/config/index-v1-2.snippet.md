---
id: vercel-server-sdk/sdk-docs/features/config/index-v1-2
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Vercel.
---

```typescript
import { createClient } from '@vercel/edge-config';
import { BasicLogger, init, LDOptions } from '@launchdarkly/vercel-server-sdk';

const edgeConfigClient = createClient(process.env.EDGE_CONFIG);

const options: LDOptions = {
  logger: new BasicLogger({ destination: console.log }),
  sendEvents: true, // default is false
};

client = init('example-client-side-id', edgeConfigClient, options);
```
