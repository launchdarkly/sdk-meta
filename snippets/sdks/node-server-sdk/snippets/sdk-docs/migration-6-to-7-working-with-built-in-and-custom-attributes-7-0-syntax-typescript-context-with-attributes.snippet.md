---
id: node-server-sdk/sdk-docs/migration-6-to-7-working-with-built-in-and-custom-attributes-7-0-syntax-typescript-context-with-attributes
sdk: node-server-sdk
kind: reference
lang: typescript
description: "7.0 syntax, Typescript: context with attributes in section \"Working with built-in and custom attributes\""
---

```ts
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
