---
id: js-client-sdk/sdk-docs/features/datasaving/standard-setup
sdk: js-client-sdk
kind: reference
lang: javascript
description: Data saving mode standard setup for JavaScript.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only
---

```js
import { createClient } from '@launchdarkly/js-client-sdk';

const client = createClient('example-client-side-id', context, {
  dataSystem: {},
});
```
