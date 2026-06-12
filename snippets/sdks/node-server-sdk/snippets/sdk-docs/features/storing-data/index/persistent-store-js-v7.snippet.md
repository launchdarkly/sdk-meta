---
id: node-server-sdk/sdk-docs/features/storing-data/index/persistent-store-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Persistent feature store configuration example for Node.js (server-side) SDK v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const ld = require('launchdarkly-node-server-sdk');

const store = SomeKindOfFeatureStore(storeOptions);
const options = {
  featureStore: store
};
const client = ld.init('YOUR_SDK_KEY', options);
```
