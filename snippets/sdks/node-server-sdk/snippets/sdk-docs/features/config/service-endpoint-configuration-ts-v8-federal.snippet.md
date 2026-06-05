---
id: node-server-sdk/sdk-docs/features/config/service-endpoint-configuration-ts-v8-federal
sdk: node-server-sdk
kind: reference
lang: typescript
description: Service endpoint configuration example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions } from '@launchdarkly/node-server-sdk';

const options: LDOptions = {
  streamUri: 'https://stream.launchdarkly.us',
  baseUri: 'https://sdk.launchdarkly.us',
  eventsUri: 'https://events.launchdarkly.us',
};
```
