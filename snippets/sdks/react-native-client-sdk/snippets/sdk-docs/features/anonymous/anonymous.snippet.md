---
id: react-native-client-sdk/sdk-docs/features/anonymous/anonymous
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: Anonymous context in a multi-context example for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```javascript
// This device context is anonymous
const deviceContext = {
  // The key attribute is required and should be empty
  // The SDK will automatically generate a unique, stable key
  key: '',
  kind: 'device',
  deviceId: '12345',
  anonymous: true
}

// This user context is not anonymous
const userContext = {
  kind: 'user',
  key: 'example-user-key'
}

// The multi-context contains one anonymous context
// and one non-anonymous context
const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}
```
