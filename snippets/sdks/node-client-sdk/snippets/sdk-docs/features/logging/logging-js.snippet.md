---
id: node-client-sdk/sdk-docs/features/logging/logging-js
sdk: node-client-sdk
kind: reference
lang: javascript
description: basicLogger debug-level configuration example for Node.js (client-side) (JavaScript).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
const LaunchDarkly = require('launchdarkly-node-client-sdk');

const options = {
  logger: LaunchDarkly.basicLogger({
    level: 'debug',
  }),
};

const client = LaunchDarkly.initialize( 'example-client-side-id', user, options);

```
