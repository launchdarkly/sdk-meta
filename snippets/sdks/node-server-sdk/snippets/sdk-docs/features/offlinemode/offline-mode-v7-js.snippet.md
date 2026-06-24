---
id: node-server-sdk/sdk-docs/features/offlinemode/offline-mode-v7-js
sdk: node-server-sdk
kind: reference
lang: javascript
description: Offline mode example for Node.js (server-side) SDK v7.x (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
const options = { offline: true };
client = ld.init('YOUR_SDK_KEY', options);
client.variation('any.feature.flag', user, false, cb) // cb will always be invoked with the default value (false)
```
