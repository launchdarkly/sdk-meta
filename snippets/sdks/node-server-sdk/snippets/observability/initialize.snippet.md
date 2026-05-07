---
id: node-server-sdk/observability/initialize
sdk: node-server-sdk
kind: initialize
lang: javascript
file: node-server-sdk/observability/initialize.txt
description: Initialize node-server-sdk with observability plugin.
---

```javascript
const ldClient = init('SDK_KEY',
  // … your existing config, if relevant
  {
  plugins: [
    new Observability({
      service: 'my-service-name',
    }),
  ],
});
```
