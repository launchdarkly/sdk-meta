---
id: react-native-client-sdk/sdk-docs/features/multienv/multiple-clients
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Multiple LDClient instances for separate environments in React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
// not recommended: support multiple environments by creating multiple clients
const client1 = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled);
const client2 = new ReactNativeLDClient('mobile-key-456def', AutoEnvAttributes.Enabled);
```
