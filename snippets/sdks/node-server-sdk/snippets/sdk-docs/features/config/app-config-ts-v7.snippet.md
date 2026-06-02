---
id: node-server-sdk/sdk-docs/features/config/app-config-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Application metadata configuration example for Node.js (server-side).
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
const options: ld.LDOptions = {
  application: {
    id: 'authentication-service',
    version: '1.0.0'
  }
};
const client = ld.init('YOUR_SDK_KEY', options);
```
