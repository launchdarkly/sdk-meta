---
id: node-server-sdk/sdk-docs/migration-6-to-7-referencing-properties-of-an-attribute-object-7-0-syntax-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, TypeScript in section \"Referencing properties of an attribute object\""
---

```ts
import * as ld from 'launchdarkly-node-server-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  address: {street: '123 Main St', city: 'Springfield'},
};
```
