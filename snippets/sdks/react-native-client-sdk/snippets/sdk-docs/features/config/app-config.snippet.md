---
id: react-native-client-sdk/sdk-docs/features/config/app-config
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: Application metadata configuration example for React Native.
---

```ts
const options: LDOptions = {
  applicationInfo: {
    id: 'authentication-service',
    name: 'Authentication-Service',
    version: '1.0.0',
    versionName: 'v1',
  }
}

const client = new ReactNativeLDClient('example-mobile-key', AutoEnvAttributes.Enabled, options);
```
