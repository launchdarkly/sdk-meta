---
id: node-server-sdk/sdk-info/flagEval-number
sdk: node-server-sdk
kind: flag-eval
lang: javascript
file: node-server-sdk/flagEval-number.txt
description: Flag evaluation example for node-server-sdk (number flag).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only
---

```javascript
// Evaluate a context
const context = {
   "kind": 'user',
   "key": 'user-key-123abc',
   "name": 'Sandy',
};

client.on('ready', () => {
  client.variation('featureKey', context, 0, function(err, flagValue) {
    if (flagValue === 1) {

      // TODO: Put your feature here

    } else {

      // TODO: Put your fallback feature here

    }
  });
});
```
