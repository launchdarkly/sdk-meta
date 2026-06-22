---
id: node-client-sdk/sdk-docs/features/privateattrs/context-js
sdk: node-client-sdk
kind: reference
lang: js
description: Marking attributes private in both the context and configuration objects for Node.js client SDK v3.0 (JavaScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```js
const context = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email']
  }
};

const client = ld.initialize('example-client-side-id', context, {
  privateAttributes: ['email']
});
```
