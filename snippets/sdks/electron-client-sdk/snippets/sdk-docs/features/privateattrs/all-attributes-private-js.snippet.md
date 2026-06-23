---
id: electron-client-sdk/sdk-docs/features/privateattrs/all-attributes-private-js
sdk: electron-client-sdk
kind: reference
lang: js
description: Marking all attributes private for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```js
const user = {
  key: 'example-user-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com'
};

const client = LDElectron.initialize('example-client-side-id', user, {
  allAttributesPrivate: true
});
```
