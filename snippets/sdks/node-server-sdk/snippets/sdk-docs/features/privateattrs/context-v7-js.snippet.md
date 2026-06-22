---
id: node-server-sdk/sdk-docs/features/privateattrs/context-v7-js
sdk: node-server-sdk
kind: reference
lang: js
description: Marking context attributes private in the context object for Node.js SDK v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
const context = {
  kind: 'user',
  key: 'example-user-key',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email'],
  }
};
```
