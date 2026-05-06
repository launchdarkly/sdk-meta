---
id: react-client-sdk/observability/initialize
sdk: react-client-sdk
kind: initialize
lang: javascript
file: react-client-sdk/observability/initialize.txt
description: Initialize react-client-sdk via withLDProvider with observability + session replay plugins.
---

```javascript
const LDProvider = withLDProvider({
  clientSideId: 'SDK_KEY',
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
