---
id: node-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-js
sdk: node-server-sdk
kind: reference
lang: javascript
description: Flag evaluation reason example for Node.js (server-side, JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
var detail = client.variationDetail('example-flag-key', context, false);

var value = detail.value;
var index = detail.variationIndex;
var reason = detail.reason;
```
