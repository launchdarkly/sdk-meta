---
id: js-client-sdk/sdk-docs/features/contextconfig/multi-context
sdk: js-client-sdk
kind: reference
lang: javascript
description: Multi-context example for JavaScript SDK v3.x+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
const deviceContext = {
  kind: 'device',
  type: 'iPad',
  key: 'example-device-key'
}

const userContext = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy',
  role: 'doctor'
}

const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}

const client = LDClient.initialize('example-client-side-id', multiContext)

try {
  await client.waitForInitialization(5);
  proceedWithSuccessfullyInitializedClient();
} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}
```
