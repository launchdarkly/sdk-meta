---
id: js-client-sdk/sdk-docs/features/securemode/configure-secure-mode-v4
sdk: js-client-sdk
kind: reference
lang: js
description: Secure mode configuration example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

// client initialization
const client = createClient('example-client-side-id', context);
client.start({ identifyOptions: { hash: 'example-server-generated-hash' } });

const result = await client.waitForInitialization({ timeout: 5 });
if (result.status === 'complete') {
  // proceed with successfully initialized client

  // identification of new contexts
  await client.identify(newContext, { hash: newContextHash });
  console.log("New context's flags available");
} else {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
