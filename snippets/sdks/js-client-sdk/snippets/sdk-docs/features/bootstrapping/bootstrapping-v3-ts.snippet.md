---
id: js-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-v3-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Bootstrapping example for JavaScript SDK v3.x (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```ts
import * as LDClient from 'launchdarkly-js-client-sdk';
import { LDContext, LDFlagSet, LDOptions } from 'launchdarkly-js-client-sdk';

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
