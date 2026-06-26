---
id: node-client-sdk/sdk-docs/features/localstorage/localstorage-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Local storage caching example for Node.js (client-side) (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
import { initialize, LDOptions } from 'launchdarkly-node-client-sdk';

const options: LDOptions = { bootstrap: 'localStorage' };

const client = initialize('example-client-side-id', context, options);
```
