---
id: node-client-sdk/sdk-docs/features/contextconfig/context-example-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Context example for Node.js (client-side) SDK v3.0 (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```typescript
import * as ld from 'launchdarkly-node-client-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  groups: ['Acme', 'Global Health Services']
};
```
