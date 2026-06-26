---
id: node-client-sdk/sdk-docs/features/localstorage/localstorage-js
sdk: node-client-sdk
kind: reference
lang: javascript
description: Local storage caching example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
const client = LDClient.initialize(
  'example-client-side-id',
  context,
  {
    bootstrap: 'localStorage'
  }
);
```
