---
id: cloudflare-server-sdk/sdk-docs/features/config/index-ts-v2-2
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Cloudflare.
---

```typescript
import { BasicLogger, init, LDOptions } from '@launchdarkly/cloudflare-server-sdk';

const options: LDOptions = {
  logger: new BasicLogger({ destination: console.log }),
};

client = init('example-client-side-id', env.LD_KV, options);
```
