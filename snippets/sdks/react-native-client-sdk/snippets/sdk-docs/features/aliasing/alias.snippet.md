---
id: react-native-client-sdk/sdk-docs/features/aliasing/alias
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: Alias event example for React Native SDK 6.x and earlier.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```js
client.alias(user, previousUser);

// to send an alias event in a different environment than the default,
// pass in the environment key (optional)
client.alias(user, previousUser, environment);
```
