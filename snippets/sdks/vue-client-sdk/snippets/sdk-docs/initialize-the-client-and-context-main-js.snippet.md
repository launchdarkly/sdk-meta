---
id: vue-client-sdk/sdk-docs/initialize-the-client-and-context-main-js
sdk: vue-client-sdk
kind: reference
lang: javascript
description: "main.js in section \"Initialize the client and context\""
---

```js
import { createApp } from 'vue'
import App from './App.vue'
import { LDPlugin } from 'launchdarkly-vue-client-sdk'
import Observability from '@launchdarkly/observability'
import SessionReplay, { LDRecord } from '@launchdarkly/session-replay'

const app = createApp(App)
app.use(LDPlugin, {
  clientSideID: 'example-client-side-id',
  deferInitialization: true,
  options: {
    plugins: [
      new Observability(),
      new SessionReplay()
    ]
  }
})
app.mount('#app')
```
