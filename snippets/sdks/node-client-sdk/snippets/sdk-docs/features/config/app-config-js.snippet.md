---
id: node-client-sdk/sdk-docs/features/config/app-config-js
sdk: node-client-sdk
kind: reference
lang: javascript
description: Application metadata configuration example for Node.js (client-side).
---

```js
const options = {
  application: {
    id: "authentication-service",
    version: "1.0.0"
  }
};

const client = LDClient.initialize('example-client-side-id', context, options);
```
