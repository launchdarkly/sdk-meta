---
id: node-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Proxy mode configuration example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const options: ld.LDOptions = {
  streamUri: 'https://your-relay-proxy.com:8030',
  baseUri: 'https://your-relay-proxy.com:8030',
  eventsUri: 'https://your-relay-proxy.com:8030',
};
```
