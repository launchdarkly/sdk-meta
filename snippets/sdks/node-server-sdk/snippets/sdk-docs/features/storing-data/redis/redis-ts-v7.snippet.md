---
id: node-server-sdk/sdk-docs/features/storing-data/redis/redis-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Redis feature store configuration example for Node.js (server-side) SDK v7.x and earlier (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions, init } from 'launchdarkly-node-server-sdk';
import { RedisFeatureStore } from 'launchdarkly-node-server-sdk-redis';

const store = RedisFeatureStore({
    redisOpts: { host: 'redis-host', port: 6379 },
    prefix: 'your-key-prefix',
    cacheTTL: 30,
});

const options: LDOptions = {
  featureStore: store,
};
const client = init('YOUR_SDK_KEY', options);
```
