---
id: node-client-sdk/sdk-docs/migration-2-to-3-referencing-properties-of-an-attribute-object-3-0-syntax-context-with-object-attributes
sdk: node-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
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
  }
};
```
