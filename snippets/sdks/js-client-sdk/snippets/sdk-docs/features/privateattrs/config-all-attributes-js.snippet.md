---
id: js-client-sdk/sdk-docs/features/privateattrs/config-all-attributes-js
sdk: js-client-sdk
kind: reference
lang: js
description: Marking all attributes private in the configuration object for JavaScript SDK v3.x+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```js

// All attributes marked private
const client = ld.initialize('example-client-side-id', context, options = {
  allAttributesPrivate: true
});

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}

```
