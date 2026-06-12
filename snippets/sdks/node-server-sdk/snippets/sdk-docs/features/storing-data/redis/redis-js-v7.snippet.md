---
id: node-server-sdk/sdk-docs/features/storing-data/redis/redis-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Redis feature store configuration example for Node.js (server-side) SDK v7.x and earlier (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('launchdarkly-node-server-sdk');
const RedisFeatureStore = require('launchdarkly-node-server-sdk-redis');

const redisOpts = {
  url: 'redis://your-redis:6379'
};
const store = RedisFeatureStore(redisOpts, 30, 'your-key-prefix');

const options = {
  featureStore: store
};
const client = ld.init(sdkKey, options);
```
