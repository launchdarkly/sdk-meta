---
id: node-server-sdk/sdk-docs/features/storing-data/redis/redis-ts-v8
sdk: node-server-sdk
kind: reference
lang: typescript
description: Redis feature store configuration example for Node.js (server-side) SDK v8.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from '@launchdarkly/node-server-sdk';
import { RedisFeatureStore } from '@launchdarkly/node-server-sdk-redis';

const store = RedisFeatureStore({
    redisOpts: { host: 'redis-host', port: 6379 },
    prefix: 'your-key-prefix',
    cacheTTL: 30,
});

const options: ld.LDOptions = {
  featureStore: store,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
