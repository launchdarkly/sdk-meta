---
id: electron-client-sdk/sdk-docs/electron-initialize-the-client-javascript-3
sdk: electron-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Initialize the client\""
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```js
// Using an event listener:
client.on('ready', function () {
  // Now we can evaluate some feature flags
})

// Or, using a Promise:
client.waitForInitialization().then(function () {
  // Now we can evaluate some feature flags
})
```
