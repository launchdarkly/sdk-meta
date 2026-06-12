---
id: react-native-client-sdk/sdk-docs/features/datasaving/standard-setup
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: Data saving mode standard setup for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```js
import { AutoEnvAttributes, ReactNativeLDClient } from '@launchdarkly/react-native-client-sdk';

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, {
  dataSystem: {},
});
```
