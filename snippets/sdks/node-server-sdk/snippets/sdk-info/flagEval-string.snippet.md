---
id: node-server-sdk/sdk-info/flagEval-string
sdk: node-server-sdk
kind: flag-eval
lang: javascript
file: node-server-sdk/flagEval-string.txt
description: Flag evaluation example for node-server-sdk (string flag).
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
  client.variation('featureKey', context, 'fallback-value', function(err, flagValue) {
    if (flagValue === 'example-variation-value') {

      // TODO: Put your feature here

    } else {

      // TODO: Put your fallback feature here

    }
  });
});
```
