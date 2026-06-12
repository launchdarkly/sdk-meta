---
id: node-server-sdk/sdk-docs/features/storing-data/consul/consul-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Consul feature store configuration example for Node.js (server-side) SDK v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
import ConsulFeatureStore from 'launchdarkly-node-server-sdk-consul';

const store = ConsulFeatureStore({
  consulOptions: {
    host: 'your-consul',
    port: 8100,
  },
  prefix: 'your-key-prefix',
  cacheTTL: 30,
});

const options: ld.LDOptions = {
  featureStore: store,
};

const client = ld.init('YOUR_SDK_KEY', options);
```
