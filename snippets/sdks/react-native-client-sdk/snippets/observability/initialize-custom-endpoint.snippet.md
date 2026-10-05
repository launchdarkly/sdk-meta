---
id: react-native-client-sdk/observability/initialize-custom-endpoint
sdk: react-native-client-sdk
kind: initialize
lang: javascript
file: react-native-client-sdk/observability/initialize-custom-endpoint.txt
description: Initialize react-native-client-sdk with the observability plugin against a custom observability domain (OBSERVABILITY_BASE_DOMAIN is substituted by the consumer).
validation:
  scaffold: react-native-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_MOBILE_KEY
---

```javascript
const client = new ReactNativeLDClient(
    'SDK_KEY',
    // … your existing config, if relevant
    AutoEnvAttributes.Enabled,
    {
      plugins: [
        new Observability({
          backendUrl: 'https://pub.OBSERVABILITY_BASE_DOMAIN',
          otlpEndpoint: 'https://otel.OBSERVABILITY_BASE_DOMAIN:4318',
          serviceName: 'example-service',
          serviceVersion: 'example-sha'
        })
      ],
    }
);
```
