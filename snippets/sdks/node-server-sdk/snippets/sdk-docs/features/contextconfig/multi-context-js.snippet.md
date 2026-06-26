---
id: node-server-sdk/sdk-docs/features/contextconfig/multi-context-js
sdk: node-server-sdk
kind: reference
lang: javascript
description: Multi-context example for Node.js (server-side) SDK v7.x and later (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
const context = {
  kind: 'multi',
  user: { key: 'example-user-key' },
  device: { key: 'example-device-key' }
}
```
