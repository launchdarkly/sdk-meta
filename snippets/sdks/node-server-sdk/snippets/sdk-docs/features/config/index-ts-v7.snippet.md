---
id: node-server-sdk/sdk-docs/features/config/index-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const options: ld.LDOptions = {
  timeout: 3,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
