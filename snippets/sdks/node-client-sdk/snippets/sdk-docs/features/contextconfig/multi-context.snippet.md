---
id: node-client-sdk/sdk-docs/features/contextconfig/multi-context
sdk: node-client-sdk
kind: reference
lang: javascript
description: Multi-context example for Node.js (client-side) SDK v3.0.
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
const deviceContext = {
  kind: 'device',
  type: 'iPad',
  key: 'example-device-key'
}

const userContext = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy',
  role: 'doctor'
}

const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}
```
