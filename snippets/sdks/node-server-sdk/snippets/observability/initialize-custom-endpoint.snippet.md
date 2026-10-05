---
id: node-server-sdk/observability/initialize-custom-endpoint
sdk: node-server-sdk
kind: initialize
lang: javascript
file: node-server-sdk/observability/initialize-custom-endpoint.txt
description: Initialize node-server-sdk with the observability plugin against a custom observability domain (OBSERVABILITY_BASE_DOMAIN is substituted by the consumer).
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
      backendUrl: 'https://pub.OBSERVABILITY_BASE_DOMAIN',
      otlpEndpoint: 'https://otel.OBSERVABILITY_BASE_DOMAIN:4318',
      service: 'my-service-name',
    }),
  ],
});
```
