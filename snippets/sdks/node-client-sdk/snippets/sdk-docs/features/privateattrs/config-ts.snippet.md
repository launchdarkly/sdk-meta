---
id: node-client-sdk/sdk-docs/features/privateattrs/config-ts
sdk: node-client-sdk
kind: reference
lang: ts
description: Private attribute configuration for Node.js client SDK v3.0 (TypeScript).
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
};
// All attributes marked private
const options: ld.LDOptions = { allAttributesPrivate: true };

const client = ld.initialize('example-client-side-id', context, options);
// Two attributes marked private
const optionsSomePrivate = { privateAttributes: ['email', 'name'] };

const clientSomePrivate = ld.initialize('example-client-side-id', context, optionsSomePrivate);
```
