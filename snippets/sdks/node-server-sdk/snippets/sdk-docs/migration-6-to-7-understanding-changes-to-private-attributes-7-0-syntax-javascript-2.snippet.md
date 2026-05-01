---
id: node-server-sdk/sdk-docs/migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-javascript-2
sdk: node-server-sdk
kind: reference
lang: javascript
description: "7.0 syntax, JavaScript in section \"Understanding changes to private attributes\""
---

```js
const context = {
  key: 'example-user-key',
  email: 'sandy@example.com',
  address: {street: '123 Main St', city: 'Springfield'},
  _meta: {
      privateAttributes: ['email', 'address']
  }
}
```
