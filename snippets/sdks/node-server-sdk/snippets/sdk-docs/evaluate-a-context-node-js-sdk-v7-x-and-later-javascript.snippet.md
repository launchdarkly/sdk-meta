---
id: node-server-sdk/sdk-docs/evaluate-a-context-node-js-sdk-v7-x-and-later-javascript
sdk: node-server-sdk
kind: reference
lang: javascript
description: "Node.js SDK v7.x and later (JavaScript) in section \"Evaluate a context\""
---

```js
const context = {
   "kind": 'user',
   "key": 'example-user-key',
   "name": 'Sandy'
};

client.on('ready', () => {
  client.variation('example-flag-key', context, false,
    (err, showFeature) => {
      if (showFeature) {
        // application code to show the feature
      } else {
        // the code to run if the feature is off
      }
    });
});
```
