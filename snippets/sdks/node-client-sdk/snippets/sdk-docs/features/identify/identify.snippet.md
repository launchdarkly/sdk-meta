---
id: node-client-sdk/sdk-docs/features/identify/identify
sdk: node-client-sdk
kind: reference
lang: javascript
description: Identify example for the Node.js client SDK v3.0 (single context).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```javascript
client.identify(newContext, null, () => {
  console.log("New context's flags available");
});

// or, with a Promise:
client.identify(newContext).then(() => {
  console.log("New context's flags available");
});
```
