---
id: node-client-sdk/sdk-docs/initialize-the-client-node-js-sdk-v4-javascript
sdk: node-client-sdk
kind: reference
lang: javascript
description: "Node.js SDK v4.x (JavaScript) in section \"Initialize the client\""
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only
---

```js
import { createClient } from '@launchdarkly/node-client-sdk';

// You'll need this context later, but you can ignore it for now.
const context = {
  kind: 'user',
  key: 'example-user-key'
};

// Create client
const client = createClient('example-client-side-id', context);

// Then start the client
client.start();

const result = await client.waitForInitialization({ timeout: 5 });

if (result.status === 'complete') {
  // initialization succeeded, flag values are now available
} else if (result.status === 'failed') {
  // initialization failed
  console.error('Initialization failed:', result.error);
} else if (result.status === 'timeout') {
  // initialization did not complete before the timeout
}
```
