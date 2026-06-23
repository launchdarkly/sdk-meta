---
id: node-client-sdk/sdk-docs/features/securemode/configure-secure-mode
sdk: node-client-sdk
kind: reference
lang: javascript
description: Secure mode configuration example for Node.js (client-side) SDK v3.
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
// client initialization
const options = {
  hash: 'example-server-generated-hash',
};
const client = LDClient.initialize('example-client-side-id', context, options);

// identification of new contexts
client.identify(newContext, hash, function() {
  console.log("New context's flags available");
});

// identification of new contexts, with a Promise
client.identify(newContext, hash).then(() => {
  console.log("New context's flags available");
});
```
