---
id: js-client-sdk/sdk-docs/features/privateattrs/config-private-attributes-v4
sdk: js-client-sdk
kind: reference
lang: js
description: Marking specific attributes private in the configuration object for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

// Two attributes marked private
const options = { privateAttributes: ['email', 'name'] };
const client = createClient('example-client-side-id', context, options);
client.start();

const result = await client.waitForInitialization({ timeout: 5 });
if (result.status === 'complete') {
  proceedWithSuccessfullyInitializedClient();
} else {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
