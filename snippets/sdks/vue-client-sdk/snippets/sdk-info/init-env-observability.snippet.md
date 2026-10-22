---
id: vue-client-sdk/sdk-info/init-env-observability
sdk: vue-client-sdk
kind: init
lang: javascript
file: vue-client-sdk/init-env-observability.txt
description: |
  Client initialization snippet for vue-client-sdk with the observability
  and session replay plugins, reading the client-side ID from an
  environment variable. Not validated, for the same reason as
  vue-client-sdk/sdk-info/init-env: the browser bundle leaves
  process.env.LAUNCHDARKLY_CLIENT_SIDE_ID undefined, and LDPlugin throws
  synchronously on a missing clientSideID during `app.use(...)`, before
  `app.mount(...)`, so even vue-syntax-only's App.vue sentinel never renders.
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

// Reads your client-side ID from the LAUNCHDARKLY_CLIENT_SIDE_ID
// environment variable, so one build can run in every environment.
// privacySetting 'strict' masks all text and inputs in session replays.
app.use(LDPlugin, {
  clientSideID: process.env.LAUNCHDARKLY_CLIENT_SIDE_ID,
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
