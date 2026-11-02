---
id: node-server-sdk/sdk-docs/features/overrides/override-affected
sdk: node-server-sdk
kind: reference
lang: javascript
description: Check whether a flag override affected an evaluation in Node.js (server-side) SDK v9.15+.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```js
const detail = await client.variationDetail('my-boolean-flag-key', context, false);
if (detail.reason.overrideAffected) {
  console.log('An override affected this evaluation');
}
```
