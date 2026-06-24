---
id: react-native-client-sdk/sdk-docs/features/flagchanges/unregister-v10
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Flag change listener unregistration for React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```typescript
const changeHandler = (context: LDContext, changedKeys: string[]) => {
  console.log('listening to change');
};
client.off('change', changeHandler);
```
