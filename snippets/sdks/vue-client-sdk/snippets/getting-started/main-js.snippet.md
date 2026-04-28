---
id: vue-client-sdk/getting-started/main-js
sdk: vue-client-sdk
kind: hello-world
lang: javascript
file: src/main.js
description: src/main.js wires the LDPlugin into the Vue app.
inputs:
  environmentId:
    type: client-side-id
    description: Client-side ID baked into the rendered source.
ld-application:
  slot: main-js
---

In `src/main.js`:

```javascript
import { createApp } from 'vue'
import App from './App.vue'
import { LDPlugin } from 'launchdarkly-vue-client-sdk'

const app = createApp(App)
app.use(LDPlugin, {
  clientSideID: '{{ environmentId }}',
  context: { kind: 'user', key: 'example-user-key' },
})
app.mount('#app')
```
