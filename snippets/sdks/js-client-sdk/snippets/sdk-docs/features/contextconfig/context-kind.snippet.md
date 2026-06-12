---
id: js-client-sdk/sdk-docs/features/contextconfig/context-kind
sdk: js-client-sdk
kind: reference
lang: javascript
description: Context with a non-user kind for JavaScript SDK v3.x+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
const context = {
  kind: 'organization',
  key: 'example-organization-key'
};
const client = LDClient.initialize('example-client-side-id', context);

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}
```
