---
id: vue-client-sdk/sdk-info/init-observability
sdk: vue-client-sdk
kind: init
lang: javascript
file: vue-client-sdk/init-observability.txt
description: Client initialization snippet for vue-client-sdk with the observability and session replay plugins.
validation:
  scaffold: vue-client-sdk/scaffolds/init-runner
  placeholders:
    YOUR_CLIENT_SIDE_ID: LAUNCHDARKLY_CLIENT_SIDE_ID
---

```javascript
// Add the code below to your main.js file.
import { createApp } from 'vue';
import App from './App.vue';
import { LDPlugin } from 'launchdarkly-vue-client-sdk';
import Observability from '@launchdarkly/observability';
import SessionReplay from '@launchdarkly/session-replay';

const app = createApp(App);

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  name: 'Sandy',
};

// This is your client-side ID.
// privacySetting 'strict' masks all text and inputs in session replays.
app.use(LDPlugin, {
  clientSideID: 'YOUR_CLIENT_SIDE_ID',
  context: context,
  options: {
    plugins: [
      new Observability(),
      new SessionReplay({ privacySetting: 'strict' }),
    ],
  },
});

app.mount('#app');
```
