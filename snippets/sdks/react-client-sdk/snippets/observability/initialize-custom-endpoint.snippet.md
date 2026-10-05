---
id: react-client-sdk/observability/initialize-custom-endpoint
sdk: react-client-sdk
kind: initialize
lang: javascript
file: react-client-sdk/observability/initialize-custom-endpoint.txt
description: Initialize react-client-sdk via withLDProvider with observability + session replay plugins against a custom observability domain (OBSERVABILITY_BASE_DOMAIN is substituted by the consumer).
validation:
  scaffold: react-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
const LDProvider = withLDProvider({
  clientSideID: 'SDK_KEY',
  // … your existing config, if relevant
  options: {
    plugins: [
      new Observability({
        backendUrl: 'https://pub.OBSERVABILITY_BASE_DOMAIN',
        otel: {
          otlpEndpoint: 'https://otel.OBSERVABILITY_BASE_DOMAIN',
        },
        networkRecording: {
          enabled: true,
          recordHeadersAndBody: true
        }
      }),
      new SessionReplay({
        // Options: 'strict', 'default', 'none'
        privacySetting: 'strict'
      })
    ]
  }
});
```
