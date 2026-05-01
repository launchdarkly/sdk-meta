---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-redis-integration-v7-x-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "v7.x (TypeScript) in section \"Understanding changes to the Redis integration\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
import { RedisFeatureStore } from 'launchdarkly-node-server-sdk-redis';

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
