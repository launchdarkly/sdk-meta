---
id: node-server-sdk/sdk-docs/features/contextconfig/context-example-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Context example for Node.js (server-side) SDK v7.x and later (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
const context = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  groups: ['Acme', 'Global Health Services'],
};
```
