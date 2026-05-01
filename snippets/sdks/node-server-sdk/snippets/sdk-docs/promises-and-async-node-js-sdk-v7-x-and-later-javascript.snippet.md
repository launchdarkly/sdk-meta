---
id: node-server-sdk/sdk-docs/promises-and-async-node-js-sdk-v7-x-and-later-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "Node.js SDK v7.x and later (JavaScript) in section \"Promises and async\""
---

```js
// Using the .then() method to add a continuation handler for a Promise
client.variation('example-flag-key', context, false).then((value) => {
  // application code
});

// Using "await" instead, within an async function
const value = await client.variation('example-flag-key', context, false);
```
