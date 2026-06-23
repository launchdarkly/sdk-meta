---
id: js-client-sdk/sdk-docs/features/privateattrs/config-private-attributes-ts
sdk: js-client-sdk
kind: reference
lang: ts
description: Marking specific attributes private in the configuration object for JavaScript SDK v3.x+ (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```ts
import * as ld from 'launchdarkly-js-client-sdk';

// Two attributes marked private
const options = { privateAttributes: ['email', 'name'] };

const client = ld.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}

```
