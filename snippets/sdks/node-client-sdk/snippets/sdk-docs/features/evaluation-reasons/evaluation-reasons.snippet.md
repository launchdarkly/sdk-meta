---
id: node-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: node-client-sdk
kind: reference
lang: javascript
description: Flag evaluation reason example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
const options = { evaluationReasons: true };
const client = LDClient.initialize('example-client-side-id', user, options);

const detail = client.variationDetail('example-flag-key', false);

const value = detail.value;
const index = detail.variationIndex;
const reason = detail.reason;
```
