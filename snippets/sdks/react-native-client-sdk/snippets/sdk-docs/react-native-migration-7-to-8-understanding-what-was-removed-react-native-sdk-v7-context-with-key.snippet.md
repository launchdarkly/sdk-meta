---
id: react-native-client-sdk/sdk-docs/react-native-migration-7-to-8-understanding-what-was-removed-react-native-sdk-v7-context-with-key
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "React Native SDK v7+, context with key in section \"Understanding what was removed\""
---

```js
const context = {
  kind: 'user',
  key: 'example-user-key'
};

const ldClient = new LDClient();
const config = {
  mobileKey: 'example-mobile-key'
};

try {
  await ldClient.configure(config, context);
} catch (err) {
  console.error(err);
}
```
