---
id: react-native-client-sdk/sdk-docs/features/hooks/define-hook
sdk: react-native-client-sdk
kind: reference
lang: ts
description: Hook implementation and configuration for the React Native SDK.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only
---

```ts
export class ExampleHook implements Hook {

  getMetadata() {
    return { name: 'Example hook'}
  }

  // Implement at least one of:
  //
  // * `beforeEvaluation` - called during the execution of a variation method
  // before the flag value has been determined
  //
  // * `afterEvaluation` - called during the execution of a variation method
  // after the flag value has been determined
  //
  // * `beforeIdentify` - called during the execution of the identify process
  // before the operation completes, but after any context modifications are performed
  //
  // * `afterIdentify` - called during the execution of the identify process
  // after the operation completes
  //
  // * `afterTrack` - called during the execution of the track process
  // after the event has been enqueued
}

const options: LDOptions = {
  hooks: [new ExampleHook()]
};

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);
```
