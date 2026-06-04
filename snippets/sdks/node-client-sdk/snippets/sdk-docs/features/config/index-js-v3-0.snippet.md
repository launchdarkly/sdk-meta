---
id: node-client-sdk/sdk-docs/features/config/index-js-v3-0
sdk: node-client-sdk
kind: reference
lang: javascript
description: SDK configuration example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```js
const options = {
  flushInterval: 10000, // milliseconds
  allAttributesPrivate: true
};

const client = LDClient.initialize('example-client-side-id', context, options);
try {
  await client.waitForInitialization(5);
  // initialization succeeded, flag values are now available
} catch (err) {
  // initialization failed or did not complete before timeout
}
```
