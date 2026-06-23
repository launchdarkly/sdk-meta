---
id: js-client-sdk/sdk-docs/features/privateattrs/context-js
sdk: js-client-sdk
kind: reference
lang: js
description: Marking specific attributes private in the context object for JavaScript SDK v3.x+.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```js
const context = {
  kind: 'user',
  key: 'example-context-key',
  name: 'Sandy Smith',
  email: 'sandy@example.com',
  _meta: {
    privateAttributes: ['email']
  }
};

```
