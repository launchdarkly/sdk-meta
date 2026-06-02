---
id: node-server-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-v8-eu
sdk: node-server-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Node.js (server-side).
---

```ts
import { LDOptions } from '@launchdarkly/node-server-sdk';

const options: LDOptions = {
  streamUri: 'https://stream.eu.launchdarkly.com',
  baseUri: 'https://sdk.eu.launchdarkly.com',
  eventsUri: 'https://events.eu.launchdarkly.com',
};
```
