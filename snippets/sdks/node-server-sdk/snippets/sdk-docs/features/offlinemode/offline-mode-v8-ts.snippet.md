---
id: node-server-sdk/sdk-docs/features/offlinemode/offline-mode-v8-ts
sdk: node-server-sdk
kind: reference
lang: typescript
description: Offline mode example for Node.js (server-side) SDK v8.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import { LDOptions, init } from '@launchdarkly/node-server-sdk';

const options: LDOptions = { offline: true };
const client = init('YOUR_SDK_KEY', options);
client.variation('any.feature.flag', context, false, cb); // cb will always be invoked with the default value (false)
```
