---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-redis-integration-v8-0-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v8.0+ (JavaScript) in section \"Understanding changes to the Redis integration\""
---

```js
const ld = require('@launchdarkly/node-server-sdk');
const RedisFeatureStore = require('@launchdarkly/node-server-sdk-redis');

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
