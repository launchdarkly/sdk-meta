---
id: node-server-sdk/sdk-docs/openfeature/evaluate-a-context
sdk: node-server-sdk
kind: reference
lang: javascript
file: node-server-sdk/sdk-docs/openfeature/evaluate-a-context.mjs
description: "Node.js OpenFeature provider in section \"Evaluate a context\""
validation:
  scaffold: node-server-sdk/scaffolds/openfeature-node-runner
---

```js
const flagValue = await client.getBooleanValue("example-flag-key", false, context);
```
