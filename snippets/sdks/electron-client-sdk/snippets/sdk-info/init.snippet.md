---
id: electron-client-sdk/sdk-info/init
sdk: electron-client-sdk
kind: init
lang: javascript
file: electron-client-sdk/init.txt
description: |
  Client initialization snippet for electron-client-sdk (main process).
  Parse/bundle-checked via the electron-syntax-only scaffold; the
  js-client image doesn't ship the electron SDK package, so the scaffold's
  dead-code guard is what keeps the require unresolved-but-parseable.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```javascript
const LDElectron = require('launchdarkly-electron-client-sdk');

// Initialize the SDK in Electron's main process. This is your
// LaunchDarkly client-side ID; it is not a secret and is safe to
// embed in client-side code.
const user = { key: 'example-user-key' };
const options = {};
const client = LDElectron.initializeInMain('YOUR_CLIENT_SIDE_ID', user, options);

client.waitForInitialization().then(function () {
  // The SDK is ready; you can evaluate flags now.
});
```
