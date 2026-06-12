---
id: node-server-sdk/sdk-docs/features/storing-data/index/persistent-store-ts-v8
sdk: node-server-sdk
kind: reference
lang: typescript
description: Persistent feature store configuration example for Node.js (server-side) SDK v8.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions, init } from '@launchdarkly/node-server-sdk';

const store = SomeKindOfFeatureStore(storeOptions);

const options: LDOptions = {
  featureStore: store,
};
const client = init('YOUR_SDK_KEY', options);
```
