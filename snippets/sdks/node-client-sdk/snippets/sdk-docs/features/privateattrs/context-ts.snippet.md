---
id: node-client-sdk/sdk-docs/features/privateattrs/context-ts
sdk: node-client-sdk
kind: reference
lang: ts
description: Marking attributes private in both the context and configuration objects for Node.js client SDK v3.0 (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```ts
import * as ld from 'launchdarkly-node-client-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email'],
  }
};
const options: ld.LDOptions = { privateAttributes: ['email'] };

const client = ld.initialize('example-client-side-id', context, options);
```
