---
id: node-server-sdk/sdk-docs/features/storing-data/consul/consul-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Consul feature store configuration example for Node.js (server-side) SDK v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('launchdarkly-node-server-sdk');
const ConsulFeatureStore = require('launchdarkly-node-server-sdk-consul');

const store = ConsulFeatureStore({
  consulOptions: {
    host: 'your-consul',
    port:  8100
  },
  prefix: 'your-key-prefix',
  cacheTTL: 30
});

const options = {
  featureStore: store
};
const client = ld.init('YOUR_SDK_KEY', options);
```
