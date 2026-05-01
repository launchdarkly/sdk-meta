---
id: js-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-anonymous-users-3-0-syntax-working-with-anonymous-contexts
sdk: js-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, working with anonymous contexts in section \"Understanding changes to anonymous users\""
---

```js
// This user context is not anonymous
const userContext = {
  kind: 'user',
  key: 'example-user-key'
}

// This device context is anonymous
// The key is omitted, and the SDK will automatically generate one
const deviceContext = {
  kind: 'device',
  deviceId: '12345',
  anonymous: true
}

// The multi-context contains one anonymous context
// and one non-anonymous context
const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}
```
