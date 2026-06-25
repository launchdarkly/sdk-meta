---
id: js-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Bootstrapping example for JavaScript SDK v4.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

const context = { kind: 'user', key: 'example-user-key'};

const client = createClient(
  'example-client-side-id', 
  context
);

// bootstrapData is the result of your server-side SDK call to get all flags
const flags = JSON.parse(bootstrapData)
const options = { bootstrap: flags }

client.start(options);
```
