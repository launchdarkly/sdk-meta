---
id: node-client-sdk/sdk-docs/features/config/index-ts-v3-0
sdk: node-client-sdk
kind: reference
lang: typescript
description: SDK configuration example for Node.js (client-side).
---

```ts
import * as LDClient from 'launchdarkly-node-client-sdk';

const options: LDClient.LDOptions = {
  flushInterval: 10000, // milliseconds
  allAttributesPrivate: true,
};
const client = LDClient.initialize('example-client-side-id', context, options);
try {
  await client.waitForInitialization(5);
  // initialization succeeded, flag values are now available
} catch (err) {
  // initialization failed or did not complete before timeout
}
```
