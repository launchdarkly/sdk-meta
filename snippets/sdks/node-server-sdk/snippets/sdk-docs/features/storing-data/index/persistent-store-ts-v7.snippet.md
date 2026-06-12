---
id: node-server-sdk/sdk-docs/features/storing-data/index/persistent-store-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Persistent feature store configuration example for Node.js (server-side) SDK v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const store = SomeKindOfFeatureStore(storeOptions);

const options: ld.LDOptions = {
  featureStore: store,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
