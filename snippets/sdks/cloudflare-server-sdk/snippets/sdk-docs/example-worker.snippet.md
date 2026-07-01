---
id: cloudflare-server-sdk/sdk-docs/example-worker
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
file: cloudflare-server-sdk/sdk-docs/example-worker.ts
description: "Cloudflare edge SDK in section \"Example Worker\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/edge-cloudflare-worker
---

```ts
import { init } from '@launchdarkly/cloudflare-server-sdk';

export default {
  async fetch(request: Request, env: Bindings): Promise<Response> {
    const context = { kind: 'user', key: 'test-user-key-1' };

    // init the ldClient, wait and finally evaluate
    const client = init('example-client-side-id', env.LD_KV);
    await client.waitForInitialization();
    const flagValue = await client.variation('flag-key', context, false);

    return new Response(`${flagValue}`);
  },
};
```
