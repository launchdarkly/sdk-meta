---
id: node-server-sdk/observability/initialize
sdk: node-server-sdk
kind: initialize
lang: javascript
file: node-server-sdk/observability/initialize.txt
description: Initialize node-server-sdk with observability plugin.
validation:
  scaffold: node-server-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_SDK_KEY
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
