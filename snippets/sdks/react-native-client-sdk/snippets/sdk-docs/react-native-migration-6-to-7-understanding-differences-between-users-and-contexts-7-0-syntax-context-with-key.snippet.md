---
id: react-native-client-sdk/sdk-docs/react-native-migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-context-with-key
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "7.0 syntax, context with key in section \"Understanding differences between users and contexts\""
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
