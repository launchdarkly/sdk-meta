---
id: node-server-sdk/sdk-docs/openfeature/construct-a-context-organization
sdk: node-server-sdk
kind: reference
lang: javascript
file: node-server-sdk/sdk-docs/openfeature/construct-a-context-organization.mjs
description: "Node.js OpenFeature provider in section \"Construct a context\" (organization)"
validation:
  scaffold: node-server-sdk/scaffolds/openfeature-node-runner
---

```js
const context = {
  kind: "organization",
  targetingKey: "example-user-key" // Could also use "key" instead of "targetingKey".
};
```
