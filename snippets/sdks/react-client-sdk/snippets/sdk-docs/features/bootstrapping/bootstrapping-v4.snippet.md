---
id: react-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-v4
sdk: react-client-sdk
kind: reference
lang: javascript
description: Bootstrapping example for React Web SDK v4.0.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

// bootstrapData is the result of your server-side SDK call to get all flags
const flags = JSON.parse(bootstrapData);

const LDProvider = createLDReactProvider(
  'example-client-side-id',
  { kind: 'user', key: 'example-user-key' },
  {
    bootstrap: flags,
  },
);
```
