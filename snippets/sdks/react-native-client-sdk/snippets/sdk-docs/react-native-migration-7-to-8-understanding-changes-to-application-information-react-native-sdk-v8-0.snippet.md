---
id: react-native-client-sdk/sdk-docs/react-native-migration-7-to-8-understanding-changes-to-application-information-react-native-sdk-v8-0
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "React Native SDK v8.0 in section \"Understanding changes to application information\""
---

```js
let client = new LDClient();

let config = {
  mobileKey: 'example-mobile-key',
  enableAutoEnvAttributes: true,
  application: {
    id: 'authentication-service',
    name: 'Authentication-Service',
    version: '1.0.0',
    versionName: 'v1',
  },
};
let context = { key: 'example-user-key', 'kind': 'user' };

await ldClient.configure(config, context);
```
