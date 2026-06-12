---
id: node-server-sdk/sdk-docs/features/storing-data/index/persistent-store-js-v8
sdk: node-server-sdk
kind: reference
lang: javascript
description: Persistent feature store configuration example for Node.js (server-side) SDK v8.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
import { init } from '@launchdarkly/node-server-sdk';

const store = SomeKindOfFeatureStore(storeOptions);
const options = {
  featureStore: store
};
const client = init('YOUR_SDK_KEY', options);
```
