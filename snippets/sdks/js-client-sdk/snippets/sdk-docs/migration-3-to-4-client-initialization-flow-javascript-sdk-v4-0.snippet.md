---
id: js-client-sdk/sdk-docs/migration-3-to-4-client-initialization-flow-javascript-sdk-v4-0
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v4.0 in section \"Client initialization flow\""
---

```js
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
