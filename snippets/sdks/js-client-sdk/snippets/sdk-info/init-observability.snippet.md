---
id: js-client-sdk/sdk-info/init-observability
sdk: js-client-sdk
kind: init
lang: javascript
file: js-client-sdk/init-observability.txt
description: Client initialization snippet for js-client-sdk with the observability and session replay plugins.
validation:
  scaffold: js-client-sdk/scaffolds/init-runner-import-lift
  placeholders:
    YOUR_CLIENT_SIDE_ID: LAUNCHDARKLY_CLIENT_SIDE_ID
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
// This is your client-side ID.
// privacySetting 'strict' masks all text and inputs in session replays.
const client = createClient('YOUR_CLIENT_SIDE_ID', context, {
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
