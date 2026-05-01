---
id: vue-client-sdk/sdk-info/init
sdk: vue-client-sdk
kind: init
lang: javascript
file: vue-client-sdk/init.txt
description: Client initialization snippet for vue-client-sdk.
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

// This is your client-side ID.
app.use(LDPlugin, {
  clientSideID: 'YOUR_CLIENT_SIDE_ID',
  context: context
});

app.mount('#app');
```
