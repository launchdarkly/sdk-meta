---
id: react-native-client-sdk/sdk-docs/react-native-migration-9-to-10-understanding-changes-to-client-initialization-client-initialization-in-react-native-sdk-v9
sdk: react-native-client-sdk
kind: reference
lang: typescript
description: "Client initialization in React Native SDK v9 in section \"Understanding changes to client initialization\""
---

```ts
let client = new LDClient();
let config: ld.LDConfig = {
  mobileKey: 'example-mobile-key',
  enableAutoEnvAttributes: true
};
let context: ld.LDContext = { key: 'example-user-key', kind: 'user' };

await client.configure(config, context);
```
