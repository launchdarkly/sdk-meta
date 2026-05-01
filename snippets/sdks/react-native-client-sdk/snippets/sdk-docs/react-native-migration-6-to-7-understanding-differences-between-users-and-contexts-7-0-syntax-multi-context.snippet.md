---
id: react-native-client-sdk/sdk-docs/react-native-migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-multi-context
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "7.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
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

await ldClient.configure(config, multiContext);
```
