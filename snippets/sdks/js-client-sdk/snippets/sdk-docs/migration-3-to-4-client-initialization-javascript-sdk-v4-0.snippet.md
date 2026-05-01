---
id: js-client-sdk/sdk-docs/migration-3-to-4-client-initialization-javascript-sdk-v4-0
sdk: js-client-sdk
kind: reference
lang: javascript
description: "JavaScript SDK v4.0 in section \"Client initialization\""
---

```js
    import { createClient } from '@launchdarkly/js-client-sdk';

    // Create client
    const client = createClient('example-client-side-id', context, options);

    // Then start the client
    client.start();
```
