---
id: vue-client-sdk/sdk-docs/configure-the-sdk-main-js
sdk: vue-client-sdk
kind: reference
lang: javascript
description: "main.js in section \"Configure the SDK\""
validation:
  scaffold: vue-client-sdk/scaffolds/vue-syntax-only
---

```js
import { createApp } from 'vue'
import App from './App.vue'
import { LDPlugin } from 'launchdarkly-vue-client-sdk'

// You'll need this context later, but you can ignore it for now.
const clientSideID = 'example-client-side-id'

const app = createApp(App)
app.use(LDPlugin, { clientSideID })
app.mount('#app')
```
