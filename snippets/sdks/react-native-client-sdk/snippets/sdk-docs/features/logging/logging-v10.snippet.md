---
id: react-native-client-sdk/sdk-docs/features/logging/logging-v10
sdk: react-native-client-sdk
kind: reference
lang: ts
description: BasicLogger debug-level configuration example for React Native SDK v10.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

---

```ts
import {
  AutoEnvAttributes,
  BasicLogger,
  type LDOptions,
  ReactNativeLDClient,
} from '@launchdarkly/react-native-client-sdk';

const options: LDOptions = {
  logger: new BasicLogger({
    level: 'debug',
    destination: console.log,
  }),
};
const featureClient = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);
```
