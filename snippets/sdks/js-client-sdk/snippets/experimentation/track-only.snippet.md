---
id: js-client-sdk/experimentation/track-only
sdk: js-client-sdk
kind: reference
lang: javascript
description: Experimentation onboarding (track only) for js-client-sdk — initialize and add a trackMetric helper for conversion events.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```javascript
import { createClient } from '@launchdarkly/js-client-sdk';

// A "context" is a data object representing users, devices, organizations, and
// other entities. You'll need this later, but you can ignore it for now.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
  email: 'biz@face.dev',
};

// This is your client-side ID.
const ldClient = createClient('YOUR_CLIENT_SIDE_ID', context);

// await start — you only need waitForInitialization() if you wait somewhere
// other than where you start the client.
await ldClient.start();
console.log('SDK successfully initialized!');

// Call trackMetric when a metric action occurs in your app —
// a click, a form submit, a page view, a custom event, whatever your metric measures.
function trackMetric(metricKey, data) {
  ldClient.track(metricKey, data);
}
```
