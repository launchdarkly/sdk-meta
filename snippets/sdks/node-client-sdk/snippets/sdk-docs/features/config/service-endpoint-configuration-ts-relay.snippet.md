---
id: node-client-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-relay
sdk: node-client-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Node.js (client-side).
---

```ts
import * as ld from 'launchdarkly-node-client-sdk';

const options: ld.LDOptions = {
  streamUrl: 'https://your-relay-proxy.com:8030',
  baseUrl: 'https://your-relay-proxy.com:8030',
  eventsUrl: 'https://your-relay-proxy.com:8030',
};
```
