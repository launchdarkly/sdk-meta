---
id: vue-client-sdk/observability/initialize
sdk: vue-client-sdk
kind: initialize
lang: javascript
file: vue-client-sdk/observability/initialize.txt
description: Initialize vue-client-sdk with observability + session replay plugins.
validation:
  scaffold: vue-client-sdk/scaffolds/init-runner-observability
  placeholders:
    SDK_KEY: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
const app = createApp(App)
app.use(LDPlugin, {
  clientSideID: 'SDK_KEY',
  deferInitialization: true,
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
})
app.mount('#app')
```
