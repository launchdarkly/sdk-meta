---
id: js-client-sdk/observability/initialize
sdk: js-client-sdk
kind: initialize
lang: javascript
file: js-client-sdk/observability/initialize.txt
description: Initialize js-client-sdk with observability + session replay plugins.
---

```javascript
const client = LDClient.initialize('SDK_KEY', {
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
