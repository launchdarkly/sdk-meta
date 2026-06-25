---
id: js-client-sdk/sdk-docs/features/localstorage/localstorage-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Local storage caching example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

const context = { kind: 'user', key: 'example-user-key'};

// Local storage is enabled by default
// You can optionally configure the maximum number of cached contexts (default is 5)
const options = { maxCachedContexts: 3 };

const client = createClient(
  'example-client-side-id',
  context,
  options
);

client.start();
```
