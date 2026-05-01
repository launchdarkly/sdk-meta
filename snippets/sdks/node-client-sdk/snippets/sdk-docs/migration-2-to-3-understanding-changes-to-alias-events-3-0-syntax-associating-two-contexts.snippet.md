---
id: node-client-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-alias-events-3-0-syntax-associating-two-contexts
sdk: node-client-sdk
kind: reference
lang: javascript
description: "3.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```js
const deviceContext = {
  kind: 'device',
  key: 'example-device-key',
  type: 'iPad'
};

const userContext = {
  kind: 'user',
  key: 'example-user-key',
  name: 'Sandy',
  role: 'doctor'
};

const multiContext = {
  kind: 'multi',
  user: userContext,
  device: deviceContext
};

client.identify(multiContext, hash, function() {
  console.log("Multi-context's flags available");
});
```
