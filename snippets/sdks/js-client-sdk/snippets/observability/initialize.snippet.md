---
id: js-client-sdk/observability/initialize
sdk: js-client-sdk
kind: initialize
lang: javascript
file: js-client-sdk/observability/initialize.txt
description: Initialize js-client-sdk with observability + session replay plugins.
validation:
  scaffold: js-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
const context = { kind: 'user', key: 'EXAMPLE_CONTEXT_KEY' };
const client = LDClient.initialize('SDK_KEY', context, {
  // … your existing config, if relevant
  plugins: [
    new Observability({
      networkRecording: {
        enabled: true,
        recordHeadersAndBody: true
      }
    }),
    new SessionReplay({
      // Options: 'strict', 'default', 'none'
      privacySetting: 'strict'
    })
  ],
});
```
