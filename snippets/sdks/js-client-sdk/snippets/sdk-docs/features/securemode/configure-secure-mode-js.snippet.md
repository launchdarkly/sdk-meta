---
id: js-client-sdk/sdk-docs/features/securemode/configure-secure-mode-js
sdk: js-client-sdk
kind: reference
lang: javascript
description: Secure mode configuration example for JavaScript SDK v3.x+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
// client initialization
const options = {
  hash: 'example-server-generated-hash',
};
const client = LDClient.initialize('example-client-side-id', context, options);

try {
  await client.waitForInitialization(5);
  // proceed with successfully initialized client

  // identification of new contexts
  client.identify(newContext, hash, function() {
    console.log("New context's flags available");
  });

} catch(err) {
  // Client failed to initialized or timed out
  // variation() calls return fallback values until initialization completes
}

```
