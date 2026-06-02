---
id: node-server-sdk/sdk-docs/features/config/app-config-js-v7
sdk: node-server-sdk
kind: reference
lang: javascript
description: Application metadata configuration example for Node.js (server-side).
---

```js
var options = {
  application: {
    id: 'authentication-service',
    version: '1.0.0'
  }
};
client = ld.init('YOUR_SDK_KEY', options);
```
