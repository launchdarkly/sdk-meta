---
id: node-client-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Proxy mode configuration example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
import * as ld from 'launchdarkly-node-client-sdk';

const options: ld.LDOptions = {
  streamUrl: 'https://your-relay-proxy.com:8030',
  baseUrl: 'https://your-relay-proxy.com:8030',
  eventsUrl: 'https://your-relay-proxy.com:8030',
};
```
