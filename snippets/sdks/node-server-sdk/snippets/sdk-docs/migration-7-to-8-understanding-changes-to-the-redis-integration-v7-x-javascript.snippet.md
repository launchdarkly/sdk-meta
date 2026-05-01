---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-the-redis-integration-v7-x-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "v7.x (JavaScript) in section \"Understanding changes to the Redis integration\""
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
