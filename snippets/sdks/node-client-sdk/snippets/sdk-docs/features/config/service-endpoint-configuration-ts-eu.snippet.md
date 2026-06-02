---
id: node-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-eu
sdk: node-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Node.js (client-side).
---

```ts
import * as ld from 'launchdarkly-node-client-sdk';

const options: ld.LDOptions = {
  streamUrl: 'https://clientstream.eu.launchdarkly.com',
  baseUrl: 'https://clientsdk.eu.launchdarkly.com',
  eventsUrl: 'https://events.eu.launchdarkly.com',
};
```
