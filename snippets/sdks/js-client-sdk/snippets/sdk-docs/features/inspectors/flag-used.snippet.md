---
id: js-client-sdk/sdk-docs/features/inspectors/flag-used
sdk: js-client-sdk
kind: reference
lang: javascript
description: Flag Used inspector configuration example for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```js
const client = LDClient.initialize(
  'example-client-side-id',
  context,
  {
    inspectors: [
      {
        type: 'flag-used',
        name: 'example-flag-used',
        method: (flagKey, flagDetail) => {
          console.log(flagKey)
          console.log(flagDetail)
        }
      }
    ]
  }
);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}
```
