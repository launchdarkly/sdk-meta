---
id: node-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-configuration-options-node-js-sdk-v8-x
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js SDK v8.x in section \"Understanding changes to configuration options\""
---

```ts
import * as ld from '@launchdarkly/node-server-sdk';

const options: ld.LDOptions = {
  proxyOptions: {
      host: 'your-proxy-host',
      port: 8080,
      scheme: 'https',
      auth: 'username:password'
    }
};
```
