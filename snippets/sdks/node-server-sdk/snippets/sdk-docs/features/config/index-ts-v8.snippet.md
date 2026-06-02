---
id: node-server-sdk/sdk-docs/features/config/index-ts-v8
sdk: node-server-sdk
kind: reference
lang: typescript
description: SDK configuration example for Node.js (server-side).
---

```ts
import * as ld from '@launchdarkly/node-server-sdk';

const options: ld.LDOptions = {
  timeout: 3,
};
const client = ld.init('YOUR_SDK_KEY', options);
```
