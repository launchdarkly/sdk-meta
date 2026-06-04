---
id: node-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-federal
sdk: node-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-client-sdk';

const options: ld.LDOptions = {
  streamUrl: 'https://clientstream.launchdarkly.us',
  baseUrl: 'https://clientsdk.launchdarkly.us',
  eventsUrl: 'https://events.launchdarkly.us',
};
```
