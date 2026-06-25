---
id: node-client-sdk/sdk-docs/features/privateattrs/config-js
sdk: node-client-sdk
kind: reference
lang: js
description: Private attribute configuration for Node.js client SDK v3.0 (JavaScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```js
const context = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com'
};
// All attributes marked private
const client = ld.initialize('example-client-side-id', context, {
  allAttributesPrivate: true
});
// Two attributes marked private
const clientSomePrivate = ld.initialize('example-client-side-id', context, {
  privateAttributes: ['email', 'name']
});
```
