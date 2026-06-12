---
id: node-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-js
sdk: node-client-sdk
kind: reference
lang: javascript
description: Bootstrapping example for Node.js (client-side) SDK v3.
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
// bootstrapData is the result of your server-side SDK call to get all flags
const flags = JSON.parse(bootstrapData)

function onPageLoad(flags) {
  // ...
  const client = LDClient.initialize(
    'example-client-side-id',
    context,
    options = {
      bootstrap: flags
    }
  );

  // ...
}
```
