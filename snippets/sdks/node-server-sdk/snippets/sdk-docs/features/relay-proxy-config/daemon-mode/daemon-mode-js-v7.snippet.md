---
id: node-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Daemon mode configuration example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const store = SomeKindOfFeatureStore(storeOptions);

const options = {
  featureStore: store,
  useLdd: true,
};
```
