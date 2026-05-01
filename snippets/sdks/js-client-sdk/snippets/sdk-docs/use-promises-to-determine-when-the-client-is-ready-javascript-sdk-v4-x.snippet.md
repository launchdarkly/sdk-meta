---
id: js-client-sdk/sdk-docs/use-promises-to-determine-when-the-client-is-ready-javascript-sdk-v4-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v4.x in section \"Use promises to determine when the client is ready\""
---

```js
  import { createClient } from '@launchdarkly/js-client-sdk';
  import { Observability } from '@launchdarkly/observability';
  import { SessionReplay } from '@launchdarkly/session-replay';

  // Create client
  const client = createClient('example-client-side-id', context, options);

  // Then start the client
  client.start();

  // Recommended: Using waitForInitialization (always resolves with status)
  const result = await client.waitForInitialization({ timeout: 5 });

  if (result.status === 'complete') {
    // Client initialized successfully
  } else if (result.status === 'failed') {
    // Client failed to initialize
    console.error('Initialization failed:', result.error);
  } else if (result.status === 'timeout') {
    // Initialization timed out
    console.error('Initialization timed out');
  }

  // Note: Events still work if you prefer that approach
  client.on('ready', () => {
    // Client is ready (success or failure)
  });
  client.on('initialized', () => {
    // Client initialized successfully
  });
  client.on('failed', (err) => {
    // Client failed to initialize
  });
```
