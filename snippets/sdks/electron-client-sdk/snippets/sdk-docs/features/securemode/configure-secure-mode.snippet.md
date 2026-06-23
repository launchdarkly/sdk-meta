---
id: electron-client-sdk/sdk-docs/features/securemode/configure-secure-mode
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Secure mode configuration example for Electron SDK.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
// client initialization
const options = {
  hash: 'example-server-generated-hash',
};
const client = LDClient.initialize('example-client-side-id', user, options);

// identification of new user
client.identify(newUser, hash, function() {
  console.log("New user's flags available");
});

// identification of new user, with a Promise
client.identify(newUser, hash).then(() => {
  console.log("New user's flags available");
});
```
