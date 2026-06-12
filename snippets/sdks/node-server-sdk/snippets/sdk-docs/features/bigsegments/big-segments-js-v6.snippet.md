---
id: node-server-sdk/sdk-docs/features/bigsegments/big-segments-js-v6
sdk: node-server-sdk
kind: reference
lang: javascript
description: Big segments Redis store configuration example for Node.js (server-side) SDK v6.2.0 - v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const LaunchDarkly = require('launchdarkly-node-server-sdk');
const { RedisBigSegmentStore } = require('launchdarkly-node-server-sdk-redis');

const store = RedisBigSegmentStore({
  redisOpts: { url: 'redis://your-redis:6379' },
  prefix: 'example-client-side-id'
});

const config = {
  bigSegments: {
    store: store,
    userCacheSize: 2000,
    userCacheTime: 30
  }
};
const client = LaunchDarkly.init(sdkKey, config);
```
