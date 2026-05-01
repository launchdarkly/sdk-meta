---
id: react-native-client-sdk/sdk-docs/react-native-migration-6-to-7-understanding-changes-to-private-attributes-7-0-syntax-attribute-marked-private-for-one-context
sdk: react-native-client-sdk
kind: reference
lang: javascript
description: "7.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```js
const context = {
  kind: 'user',
  key: 'example-user-key',
  firstName: 'Sandy',
  lastName: 'Smith',
  email: 'sandy@example.com',
  address: {
    street: '123 Main St',
    city: 'Springfield'
  },
  _meta: {
    privateAttributes: ['email', '/address/street']
  }
};
```
