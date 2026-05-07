---
id: node-client-sdk/sdk-docs/initialize-the-client-javascript
sdk: node-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Initialize the client\""
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
client.on('initialized', () => {
  // initialization succeeded, flag values are now available
  const flagValue = client.variation('example-flag-key', false);
  // etc.
});
```
