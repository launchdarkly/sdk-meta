---
id: node-server-sdk/sdk-docs/features/logging/logging-v7-ts
sdk: node-server-sdk
kind: reference
lang: ts
description: basicLogger debug-level configuration example for Node.js SDK v7.x and earlier (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const logger: ld.LDLogger = ld.basicLogger({
  level: 'debug',
  destination: console.log,
});

const options: ld.LDOptions = { logger: logger };
```
