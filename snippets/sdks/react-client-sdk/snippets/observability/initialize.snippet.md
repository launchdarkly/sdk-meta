---
id: react-client-sdk/observability/initialize
sdk: react-client-sdk
kind: initialize
lang: javascript
file: react-client-sdk/observability/initialize.txt
description: Initialize react-client-sdk via withLDProvider with observability + session replay plugins.
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
