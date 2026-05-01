---
id: node-server-sdk/sdk-info/flagEval
sdk: node-server-sdk
kind: flag-eval
lang: javascript
file: node-server-sdk/flagEval.txt
description: Flag evaluation example for node-server-sdk.
---

```javascript
// Evaluate a context
const context = {
   "kind": 'user',
   "key": 'user-key-123abc',
   "name": 'Sandy',
};

client.on('ready', () => {
  client.variation('featureKey', context, false, function(err, showFeature) {
    if (showFeature) {

      // TODO: Put your feature here

    } else {

      // TODO: Put your fallback feature here

    }
  });
});
```
