---
id: js-client-sdk/sdk-info/init-env-observability
sdk: js-client-sdk
kind: init
lang: javascript
file: js-client-sdk/init-env-observability.txt
description: Client initialization snippet for js-client-sdk with the observability and session replay plugins, reading the client-side ID from an environment variable.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
import Observability from '@launchdarkly/observability';
import SessionReplay from '@launchdarkly/session-replay';

// A "context" is a data object representing users, devices, organizations, and
// other entities. You'll need this later, but you can ignore it for now.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY'
};
// Reads your client-side ID from the LAUNCHDARKLY_CLIENT_SIDE_ID
// environment variable, so one build can run in every environment.
// privacySetting 'strict' masks all text and inputs in session replays.
const client = createClient(process.env.LAUNCHDARKLY_CLIENT_SIDE_ID, context, {
  plugins: [
    new Observability(),
    new SessionReplay({ privacySetting: 'strict' }),
  ],
});
client.start();

const { status } = await client.waitForInitialization();

if (status === 'complete') {
  console.log('SDK successfully initialized!');
} else {
  console.error('Initialization failed');
}
```
