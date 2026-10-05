---
id: react-client-sdk/observability/initialize-no-session-replay
sdk: react-client-sdk
kind: initialize
lang: javascript
file: react-client-sdk/observability/initialize-no-session-replay.txt
description: Initialize react-client-sdk via withLDProvider with the observability plugin only (no session replay).
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
      })
    ]
  }
});
```
