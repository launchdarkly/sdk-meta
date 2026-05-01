---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-typescript-2
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, TypeScript in section \"Understanding changes to private attributes\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';
const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  email: 'sandy@example.com',
  address: {street: '123 Main St', city: 'Springfield'},
  _meta: {
    privateAttributes: ['email', 'address'],
   }
 }
```
