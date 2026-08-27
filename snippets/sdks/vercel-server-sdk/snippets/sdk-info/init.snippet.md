---
id: vercel-server-sdk/sdk-info/init
sdk: vercel-server-sdk
kind: init
lang: typescript
file: vercel-server-sdk/init.txt
description: |
  Client initialization snippet for vercel-server-sdk. Type-checked
  against the real packages via the edge-vercel-toplevel scaffold
  (runtime validation would need a live Vercel Edge Config populated by
  the LaunchDarkly integration).
validation:
  scaffold: vercel-server-sdk/scaffolds/edge-vercel-toplevel
---

```typescript
import { init } from '@launchdarkly/vercel-server-sdk';
import { createClient } from '@vercel/edge-config';

// Create an Edge Config client, then initialize the LaunchDarkly client
// with your client-side ID.
const edgeConfigClient = createClient(process.env.EDGE_CONFIG);
const ldClient = init('YOUR_CLIENT_SIDE_ID', edgeConfigClient);

await ldClient.waitForInitialization();
```
