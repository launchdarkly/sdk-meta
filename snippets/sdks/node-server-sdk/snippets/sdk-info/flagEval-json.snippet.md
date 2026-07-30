---
id: node-server-sdk/sdk-info/flagEval-json
sdk: node-server-sdk
kind: flag-eval
lang: javascript
file: node-server-sdk/flagEval-json.txt
description: Flag evaluation example for node-server-sdk (JSON flag).
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
  client.variation('featureKey', context, {}, function(err, flagValue) {
    // Use the flag value to configure your feature
    // TODO: Put your feature here
  });
});
```
