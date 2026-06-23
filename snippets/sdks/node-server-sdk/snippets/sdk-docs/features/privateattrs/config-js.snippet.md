---
id: node-server-sdk/sdk-docs/features/privateattrs/config-js
sdk: node-server-sdk
kind: reference
lang: js
description: Private attribute configuration for Node.js SDK v7.0+ (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```js
// All attributes marked private
const options = {
  allAttributesPrivate: true
};
client = ld.init('YOUR_SDK_KEY', options);

// Two attributes marked private
const optionsSomePrivate = {
  privateAttributes: ['email', 'address']
};
client = ld.init('YOUR_SDK_KEY', optionsSomePrivate);
```
