---
id: node-client-sdk/sdk-docs/initialize-the-client-node-js-sdk-v3-javascript
sdk: node-client-sdk
kind: reference
lang: javascript
description: "Node.js SDK v3 (JavaScript) in section \"Initialize the client\""
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
// You'll need this context later, but you can ignore it for now.
const context = {
  kind: 'user',
  key: 'example-user-key'
};

const client = LaunchDarkly.initialize('example-client-side-id', context);
try {
  await client.waitForInitialization(5);
  // initialization succeeded, flag values are now available
} catch (err) {
  // initialization failed or did not complete before timeout
}
```
