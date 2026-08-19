---
id: vue-client-sdk/sdk-info/init-env
sdk: vue-client-sdk
kind: init
lang: javascript
file: vue-client-sdk/init-env.txt
description: Client initialization snippet for vue-client-sdk reading the client-side ID from an environment variable.
validation:
  scaffold: vue-client-sdk/scaffolds/vue-syntax-only
---

```javascript
// Add the code below to your main.js file.
import { createApp } from 'vue';
import App from './App.vue';
import { LDPlugin } from 'launchdarkly-vue-client-sdk';

const app = createApp(App);

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  name: 'Sandy',
};

// Reads your client-side ID from the LAUNCHDARKLY_CLIENT_SIDE_ID
// environment variable, so one build can run in every environment.
app.use(LDPlugin, {
  clientSideID: process.env.LAUNCHDARKLY_CLIENT_SIDE_ID,
  context: context
});

app.mount('#app');
```
