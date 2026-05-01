---
id: js-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-private-attributes-3-0-syntax-attributes-marked-private-for-particular-context
sdk: js-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, attributes marked private for particular context in section \"Understanding changes to private attributes\""
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
