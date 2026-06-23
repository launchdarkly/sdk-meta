---
id: node-client-sdk/sdk-docs/features/logging/logging-ts
sdk: node-client-sdk
kind: reference
lang: ts
description: basicLogger debug-level configuration example for Node.js (client-side) (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```ts
import * as LaunchDarkly from 'launchdarkly-node-client-sdk';

const logger: LaunchDarkly.LDLogger = LaunchDarkly.basicLogger({ level: 'debug' });
const options: LaunchDarkly.LDOptions = { logger: logger };

const client = LaunchDarkly.initialize( 'example-client-side-id', user, options);

```
