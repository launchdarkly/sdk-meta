---
id: react-native-client-sdk/sdk-docs/react-native-migration-6-to-7-understanding-changes-to-alias-events-3-0-syntax-identifying-a-multi-context
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, identifying a multi-context in section \"Understanding changes to alias events\""
---

```js
const deviceContext = {
  kind: 'device',
  key: 'example-device-key',
  type: 'iPad'
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

try {
  await client.identify(multiContext);
  console.log("Multi-context's flags available");
} catch (err) {
  console.error(err);
}
```
