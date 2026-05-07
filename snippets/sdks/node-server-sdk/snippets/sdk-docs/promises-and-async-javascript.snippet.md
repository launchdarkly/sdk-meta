---
id: node-server-sdk/sdk-docs/promises-and-async-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Promises and async\""
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
// Using .then() and .catch() to add success and error handlers to a Promise
client.waitForInitialization({timeout: 10}).then((client) => {
  // initialization complete
}).catch((err) => {
  // timeout or initialization failed
});

// Using "await" instead, within an async function
try {
  await client.waitForInitialization({timeout: 10});
  // initialization complete
} catch (err) {
  // timeout or initialization failed
}
```
