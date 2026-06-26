---
id: node-server-sdk/sdk-docs/features/bigsegments/big-segments-ts-v6
sdk: node-server-sdk
kind: reference
lang: typescript
description: Big segments Redis store configuration example for Node.js (server-side) SDK v6.2.0 - v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as LaunchDarkly from 'launchdarkly-node-server-sdk';
import { RedisBigSegmentStore } from 'launchdarkly-node-server-sdk-redis';

const store = RedisBigSegmentStore({
  redisOpts: { url: 'redis://your-redis:6379' },
  prefix: 'example-client-side-id'
});

const options: LaunchDarkly.LDOptions = {
  bigSegments: {
    store: store,
    userCacheSize: 2000,
    userCacheTime: 30
  }
};
const client = LaunchDarkly.init('YOUR_SDK_KEY', options);
```
