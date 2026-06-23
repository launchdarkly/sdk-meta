---
id: js-client-sdk/sdk-docs/features/privateattrs/config-all-attributes-ts
sdk: js-client-sdk
kind: reference
lang: ts
description: Marking all attributes private in the configuration object for JavaScript SDK v3.x+ (TypeScript).
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```ts
import * as ld from 'launchdarkly-js-client-sdk';

// All attributes marked private
const options: ld.LDOptions = { allAttributesPrivate: true };

const client = ld.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}
```
