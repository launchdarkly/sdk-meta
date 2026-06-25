---
id: js-client-sdk/sdk-docs/features/localstorage/localstorage-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: Local storage caching example for JavaScript SDK v3.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
const options = { bootstrap: 'localStorage' }

const client = LDClient.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
