---
id: node-server-sdk/sdk-docs/features/bigsegments/big-segments-js-v8
sdk: node-server-sdk
kind: reference
lang: javascript
description: Big segments Redis store configuration example for Node.js (server-side) SDK v8.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
import { init } from '@launchdarkly/node-server-sdk';
import { RedisBigSegmentStore } from '@launchdarkly/node-server-sdk-redis';

const store = RedisBigSegmentStore({
  redisOpts: {
    host: 'your-redis',
    port: 6379
  },
  prefix: 'example-client-side-id'
});

const config = {
  bigSegments: {
    store: store,
    userCacheSize: 2000,
    userCacheTime: 30
  }
};
const client = init("YOUR_SDK_KEY", config);
```
