---
id: node-client-sdk/sdk-docs/features/identify/identify-multi
sdk: node-client-sdk
kind: reference
lang: javascript
description: Identify example for the Node.js client SDK v3.0 (multi-context).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```javascript
const newMultiContext = {
  kind: 'multi',
  user: {
    key: 'aa0ceb',
    name: 'Anna',
    role: 'doctor'
  },
  device: {
    key: 'example-device-key',
    platform: 'Android'
  }
}

client.identify(newMultiContext, null, () => {
  console.log("New context's flags available");
});

// or, with a Promise:
client.identify(newMultiContext).then(() => {
  console.log("New context's flags available");
});
```
