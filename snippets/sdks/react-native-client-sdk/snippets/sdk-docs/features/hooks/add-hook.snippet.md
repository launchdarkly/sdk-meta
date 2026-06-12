---
id: react-native-client-sdk/sdk-docs/features/hooks/add-hook
sdk: react-native-client-sdk
kind: reference
lang: ts
description: Adding a hook to an existing client for the React Native SDK.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```ts

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);
client.addHook(new ExampleHook());
```
