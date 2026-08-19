---
id: node-client-sdk/sdk-info/init-env
sdk: node-client-sdk
kind: init
lang: javascript
file: node-client-sdk/init-env.txt
description: Client initialization snippet for node-client-sdk reading the client-side ID from an environment variable.
validation:
  scaffold: node-client-sdk/scaffolds/init-runner
---

```javascript
import * as LaunchDarkly from 'launchdarkly-node-client-sdk';

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
};

// Reads your client-side ID from the LAUNCHDARKLY_CLIENT_SIDE_ID
// environment variable, so one build can run in every environment.
const client = LaunchDarkly.initialize(process.env.LAUNCHDARKLY_CLIENT_SIDE_ID, context);

try {
  await client.waitForInitialization(5);
  // Initialization succeeded
} catch (err) {
  // Initialization failed or did not complete before timeout
}
```
