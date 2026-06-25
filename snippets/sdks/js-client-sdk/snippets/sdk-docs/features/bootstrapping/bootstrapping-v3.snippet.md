---
id: js-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: Bootstrapping example for JavaScript SDK v3.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
// bootstrapData is the result of your server-side SDK call to get all flags
const flags = JSON.parse(bootstrapData)

function onPageLoad(flags) {
  // ...
  const options = { bootstrap: flags };
  const client = LDClient.initialize(
    'example-client-side-id',
    context,
    options
  );

  // ...
}
```
