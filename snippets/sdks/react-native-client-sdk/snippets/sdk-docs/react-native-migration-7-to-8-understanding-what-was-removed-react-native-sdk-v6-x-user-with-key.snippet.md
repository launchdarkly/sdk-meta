---
id: react-native-client-sdk/sdk-docs/react-native-migration-7-to-8-understanding-what-was-removed-react-native-sdk-v6-x-user-with-key
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "React Native SDK v6.x, user with key in section \"Understanding what was removed\""
---

```js
const user = {
  key: 'example-user-key'
};

const ldClient = new LDClient();
const config = {
  mobileKey: 'example-mobile-key'
};

try {
  await ldClient.configure(config, user);
} catch (err) {
  console.error(err);
}
```
