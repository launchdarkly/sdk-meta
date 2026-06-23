---
id: react-native-client-sdk/sdk-docs/features/privateattrs/config-v10
sdk: react-native-client-sdk
kind: reference
lang: js
description: Private attribute configuration for React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```js
// All attributes marked private
const options = {
  allAttributesPrivate: true
}
const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);

// Two attributes marked private
const optionsSomePrivate = {
  privateAttributes: ['email', 'address']
}
const clientSomePrivate = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, optionsSomePrivate);
```
