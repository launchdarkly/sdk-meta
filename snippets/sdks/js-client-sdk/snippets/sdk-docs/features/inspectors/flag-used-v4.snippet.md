---
id: js-client-sdk/sdk-docs/features/inspectors/flag-used-v4
sdk: js-client-sdk
kind: reference
lang: js
description: Flag Used inspector configuration example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

const client = createClient(
  'example-client-side-id',
  context,
  {
    inspectors: [
      {
        type: 'flag-used',
        name: 'example-flag-used',
        method: (flagKey, flagDetail) => {
          console.log(flagKey)
          console.log(flagDetail)
        }
      }
    ]
  }
);
client.start();

const result = await client.waitForInitialization({ timeout: 5 });
if (result.status === 'complete') {
  proceedWithSuccessfullyInitializedClient();
} else {
  // Client failed to initialize or timed out
  // variation() calls return fallback values until initialization completes
}
```
