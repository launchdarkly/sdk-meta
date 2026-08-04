---
id: node-client-sdk/sdk-docs/initialize-the-client-v4-javascript
sdk: node-client-sdk
kind: reference
lang: javascript
description: "JavaScript, Node.js SDK v4.x in section \"Initialize the client\" (event listeners)"
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
client.on('ready', () => {
  // The client has finished starting up, whether or not it succeeded
});
client.on('initialized', () => {
  // Initialization succeeded, flag values are now available
  const flagValue = client.variation('example-flag-key', false);
  // etc.
});
```
