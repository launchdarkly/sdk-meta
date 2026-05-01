---
id: js-client-sdk/sdk-docs/migration-2-to-3-understanding-differences-between-users-and-contexts-3-0-syntax-multi-context
sdk: js-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```js
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

const client = LDClient.initialize('example-client-side-id', multiContext)
```
