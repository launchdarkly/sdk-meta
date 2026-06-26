---
id: node-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Bootstrapping example for Node.js (client-side) SDK v3 (TypeScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```ts
import * as LDClient from 'launchdarkly-node-client-sdk';
import { LDContext, LDFlagSet, LDOptions } from 'launchdarkly-node-client-sdk';

// bootstrapData is the result of your server-side SDK call to get all flags
const flags: LDFlagSet = JSON.parse(bootstrapData)

function onPageLoad(flags: LDFlagSet) {
  // ...
  const context: LDContext = { kind: 'user', key: 'example-user-key'}
  const options: LDOptions = { bootstrap: flags };
  const client = LDClient.initialize('example-client-side-id', context, options);
  // ...
}
```
