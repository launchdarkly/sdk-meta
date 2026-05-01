---
id: react-native-client-sdk/sdk-docs/react-native-migration-6-to-7-understanding-differences-between-users-and-contexts-7-0-syntax-single-context-with-key
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "7.0 syntax, single context with key in section \"Understanding differences between users and contexts\""
---

```js
const context = {
  kind: 'organization',
  key: 'example-organization-key'
};

await ldClient.configure(config, context);
```
