---
id: node-client-sdk/sdk-docs/features/config/app-config-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Application metadata configuration example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```ts
import * as LDClient from 'launchdarkly-node-client-sdk';

const options: LDClient.LDOptions = {
  application: {
    id: "authentication-service",
    version: "1.0.0"
  }
};
const client = LDClient.initialize('example-client-side-id', context, options);
```
