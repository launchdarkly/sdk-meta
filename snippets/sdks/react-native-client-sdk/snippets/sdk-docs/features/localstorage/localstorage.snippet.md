---
id: react-native-client-sdk/sdk-docs/features/localstorage/localstorage
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Local storage caching example for React Native.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```ts
import { ReactNativeLDClient, AutoEnvAttributes } from '@launchdarkly/react-native-client-sdk';

// Local storage is enabled by default using AsyncStorage
// You can optionally configure the maximum number of cached contexts (default is 5)
const options = { maxCachedContexts: 3 };

const client = new ReactNativeLDClient(
  'example-mobile-key',
  AutoEnvAttributes.Enabled,
  options
);

await client.start();
```
