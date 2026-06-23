---
id: node-server-sdk/sdk-docs/features/privateattrs/context-v7-ts
sdk: node-server-sdk
kind: reference
lang: ts
description: Marking context attributes private in the context object for Node.js SDK v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const user: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email'],
  }
};
```
