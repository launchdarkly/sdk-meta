---
id: node-server-sdk/sdk-docs/features/privateattrs/context-ai-js
sdk: node-server-sdk
kind: reference
lang: js
description: Marking context attributes private for the Node.js (server-side) AI SDK (JavaScript).
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
