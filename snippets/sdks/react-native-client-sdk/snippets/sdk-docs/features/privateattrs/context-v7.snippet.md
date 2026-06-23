---
id: react-native-client-sdk/sdk-docs/features/privateattrs/context-v7
sdk: react-native-client-sdk
kind: reference
lang: js
description: Marking context attributes private for React Native SDK v7+.
validation:
  scaffold: react-native-client-sdk/scaffolds/react-native-syntax-only

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
