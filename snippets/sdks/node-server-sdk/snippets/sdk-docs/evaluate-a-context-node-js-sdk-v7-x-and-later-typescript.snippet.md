---
id: node-server-sdk/sdk-docs/evaluate-a-context-node-js-sdk-v7-x-and-later-typescript
sdk: node-server-sdk
kind: reference
lang: typescript
description: "Node.js SDK v7.x and later (TypeScript) in section \"Evaluate a context\""
---

```ts
const context = {
   "kind": 'user',
   "key": 'example-user-key',
   "name": 'Sandy',
};

client.on('ready', () => {
  client.variation('example-flag-key', context, false, function(err, showFeature) {
    client.track('event-called', context);
    if (showFeature) {
      // application code to show the feature
    } else {
      // the code to run if the feature is off
    }
  });
});
```
