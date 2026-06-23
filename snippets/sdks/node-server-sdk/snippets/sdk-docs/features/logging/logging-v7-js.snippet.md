---
id: node-server-sdk/sdk-docs/features/logging/logging-v7-js
sdk: node-server-sdk
kind: reference
lang: javascript
description: basicLogger debug-level configuration example for Node.js SDK v7.x and earlier (JavaScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
const ld = require('launchdarkly-node-server-sdk');

const options = {
  logger: ld.basicLogger({
    level: 'debug',
  }),
};
```
