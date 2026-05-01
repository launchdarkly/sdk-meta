---
id: node-client-sdk/sdk-info/init
sdk: node-client-sdk
kind: init
lang: javascript
file: node-client-sdk/init.txt
description: Client initialization snippet for node-client-sdk.
---

```javascript
import * as LaunchDarkly from 'launchdarkly-node-client-sdk';

// A "context" is a data object representing users, devices, organizations, and other entities.
const context = {
  kind: 'user',
  key: 'EXAMPLE_CONTEXT_KEY',
};

// This is your client-side ID.
const client = LaunchDarkly.initialize('YOUR_CLIENT_SIDE_ID', context);

try {
  await client.waitForInitialization(5);
  // Initialization succeeded
} catch (err) {
  // Initialization failed or did not complete before timeout
}
```
