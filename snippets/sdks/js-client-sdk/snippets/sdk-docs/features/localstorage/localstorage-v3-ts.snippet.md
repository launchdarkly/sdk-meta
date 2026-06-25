---
id: js-client-sdk/sdk-docs/features/localstorage/localstorage-v3-ts
sdk: js-client-sdk
kind: reference
lang: typescript
description: Local storage caching example for JavaScript SDK v3.x (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```ts
import * as LDClient from 'launchdarkly-js-client-sdk';
import { LDOptions } from 'launchdarkly-js-client-sdk';

const options: LDOptions = { bootstrap: 'localStorage' };
const client = LDClient.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
