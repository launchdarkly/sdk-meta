---
id: node-server-sdk/sdk-docs/features/contextconfig/context-example-ts-v7
sdk: node-server-sdk
kind: reference
lang: typescript
description: Context example for Node.js (server-side) SDK v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import * as ld from 'launchdarkly-node-server-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  groups: ['Acme', 'Global Health Services'],
};
```
