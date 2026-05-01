---
id: js-client-sdk/sdk-docs/migration-3-to-4-client-initialization-flow-javascript-sdk-v3-x
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v3.x in section \"Client initialization flow\""
---

```js
    // Option 1: Using waitUntilReady (never rejects)
    await client.waitUntilReady();

    // Option 2: Using waitForInitialization (rejects on failure)
    try {
    await client.waitForInitialization(5);
    } catch (err) {
    // Failure - but this could be an unhandled rejection if not caught
    }

    // Option 3: Using event listeners
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
