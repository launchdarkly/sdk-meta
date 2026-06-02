---
id: react-native-client-sdk/sdk-docs/features/config/index
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: SDK configuration example for React Native.
---

```ts
import { type LDOptions } from '@launchdarkly/react-native-client-sdk';

const options = {
  withReasons: true,
};

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);
```
