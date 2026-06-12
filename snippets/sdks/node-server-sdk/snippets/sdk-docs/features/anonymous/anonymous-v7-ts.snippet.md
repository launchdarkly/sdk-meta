---
id: node-server-sdk/sdk-docs/features/anonymous/anonymous-v7-ts
sdk: node-server-sdk
kind: reference
lang: typescript
description: Anonymous context example for Node.js (server-side), SDK v7.x (TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```typescript
import * as ld from 'launchdarkly-node-server-sdk';

const context: ld.LDContext = {
  kind: 'user',
  key: 'example-user-key',
  anonymous: true,
}
```
