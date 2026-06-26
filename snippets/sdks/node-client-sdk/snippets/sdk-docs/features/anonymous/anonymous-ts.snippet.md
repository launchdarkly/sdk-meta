---
id: node-client-sdk/sdk-docs/features/anonymous/anonymous-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Anonymous context example for Node.js (client-side) (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```typescript
import * as ld from 'launchdarkly-node-client-sdk';

const anonymousContext: ld.LDContext = { kind: 'user', key: 'example-user-key', anonymous: true };
```
