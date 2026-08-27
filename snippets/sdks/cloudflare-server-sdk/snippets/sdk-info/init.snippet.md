---
id: cloudflare-server-sdk/sdk-info/init
sdk: cloudflare-server-sdk
kind: init
lang: typescript
file: cloudflare-server-sdk/init.txt
description: |
  Client initialization snippet for cloudflare-server-sdk: the full
  Worker shape from the docs, since init only makes sense inside a fetch
  handler with the LD_KV binding. Type-checked against the real package
  via the edge-cloudflare-worker scaffold (runtime validation would need
  a live Workers KV namespace).
validation:
  scaffold: cloudflare-server-sdk/scaffolds/edge-cloudflare-worker
---

```typescript
import { init } from '@launchdarkly/cloudflare-server-sdk';

export default {
  async fetch(request: Request, env: Bindings): Promise<Response> {
    const context = { kind: 'user', key: 'example-user-key' };

    // Initialize the client with your LaunchDarkly client-side ID and
    // the KV namespace bound by the LaunchDarkly Cloudflare integration.
    const client = init('YOUR_CLIENT_SIDE_ID', env.LD_KV);
    await client.waitForInitialization();
    const flagValue = await client.variation('sample-feature', context, false);

    return new Response(`sample-feature: ${flagValue}`);
  },
};
```
