---
id: node-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Daemon mode configuration example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```ts
import { LDOptions } from 'launchdarkly-node-server-sdk';

const store = SomeKindOfFeatureStore(storeOptions);

const options: LDOptions = {
  featureStore: store,
  useLdd: true,
};
```
