---
id: react-native-client-sdk/sdk-docs/features/contextconfig/multi-context
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Multi-context example for React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
const deviceContext = {
  kind: 'device',
  key: 'example-device-key'
};

const userContext = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy',
  role: 'doctor'
};

const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
}
```
