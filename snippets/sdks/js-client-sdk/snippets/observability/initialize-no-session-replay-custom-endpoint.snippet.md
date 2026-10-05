---
id: js-client-sdk/observability/initialize-no-session-replay-custom-endpoint
sdk: js-client-sdk
kind: initialize
lang: javascript
file: js-client-sdk/observability/initialize-no-session-replay-custom-endpoint.txt
description: Initialize js-client-sdk with the observability plugin only against a custom observability domain (OBSERVABILITY_BASE_DOMAIN is substituted by the consumer).
validation:
  scaffold: js-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
const context = { kind: 'user', key: 'EXAMPLE_CONTEXT_KEY' };
const client = createClient('SDK_KEY', context, {
  // … your existing config, if relevant
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
    })
  ],
});
client.start();
```
