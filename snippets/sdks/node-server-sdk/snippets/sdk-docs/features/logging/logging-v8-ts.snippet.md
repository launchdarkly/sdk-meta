---
id: node-server-sdk/sdk-docs/features/logging/logging-v8-ts
sdk: node-server-sdk
kind: reference
lang: ts
description: basicLogger debug-level configuration example for Node.js SDK v8.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import { LDOptions, LDLogger, basicLogger } from '@launchdarkly/node-server-sdk';

const logger: LDLogger = basicLogger({
  level: 'debug',
  destination: console.log,
});

const options: LDOptions = { logger: logger };
```
