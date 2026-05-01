---
id: node-client-sdk/sdk-docs/initialize-the-client-javascript
sdk: node-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Initialize the client\""
---

```js
client.on('initialized', () => {
  // initialization succeeded, flag values are now available
  const flagValue = client.variation('example-flag-key', false);
  // etc.
});
```
