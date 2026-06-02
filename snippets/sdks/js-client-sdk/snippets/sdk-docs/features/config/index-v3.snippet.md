---
id: js-client-sdk/sdk-docs/features/config/index-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: SDK configuration example for JavaScript.
---

```js
const options = { allAttributesPrivate: true };
const client = LDClient.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
