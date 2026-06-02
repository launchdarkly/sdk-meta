---
id: js-client-sdk/sdk-docs/features/config/index-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: SDK configuration example for JavaScript.
---

```js
const options = { allAttributesPrivate: true };
const client = createClient('example-client-side-id', context, options);
client.start();

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
```
