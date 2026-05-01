---
id: vue-client-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-multi-context
sdk: vue-client-sdk
kind: reference
lang: javascript
description: "Multi-context in section \"Understanding differences between users and contexts\""
---

```js
const deviceContext = { kind: 'device', key: 'example-device-key', type: 'iPad' }
const userContext = { kind: 'user', key: 'example-user-key', name: 'Sandy' }

const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}
```
