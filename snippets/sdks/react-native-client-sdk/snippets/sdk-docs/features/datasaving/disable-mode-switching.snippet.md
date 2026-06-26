---
id: react-native-client-sdk/sdk-docs/features/datasaving/disable-mode-switching
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: Disable automatic mode switching for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```js
const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, {
  dataSystem: {
    automaticModeSwitching: false,
  },
});
```
