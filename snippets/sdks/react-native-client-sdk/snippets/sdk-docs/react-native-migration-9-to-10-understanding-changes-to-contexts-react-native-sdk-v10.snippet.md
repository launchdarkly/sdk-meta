---
id: react-native-client-sdk/sdk-docs/react-native-migration-9-to-10-understanding-changes-to-contexts-react-native-sdk-v10
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "React Native SDK v10 in section \"Understanding changes to contexts\""
---

```js
// This device context is anonymous
const deviceContext = {
  // The key attribute is required and should be empty
  // The SDK will automatically generate a unique, stable key
  key: '',
  kind: 'device',
  deviceId: '12345',
  anonymous: true
}
```
