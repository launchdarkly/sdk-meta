---
id: node-server-sdk/sdk-docs/features/storing-data/redis/redis-js-v8
sdk: node-server-sdk
kind: reference
lang: javascript
description: Redis feature store configuration example for Node.js (server-side) SDK v8.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('@launchdarkly/node-server-sdk');
const { RedisFeatureStore } = require('@launchdarkly/node-server-sdk-redis');

const store = RedisFeatureStore({
    redisOpts: { host: 'redis-host', port: 6379 },
    prefix: 'your-key-prefix',
    cacheTTL: 30,
});

const options = {
  featureStore: store
};
const client = ld.init(sdkKey, options);
```
