---
id: react-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-v4-start-options
sdk: react-client-sdk
kind: reference
lang: javascript
description: Bootstrapping through startOptions for React Web SDK v4.0.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

const LDProvider = createLDReactProvider(
  'example-client-side-id',
  { kind: 'user', key: 'example-user-key' },
  {
    startOptions: {
      bootstrap: flags,
    },
  },
);
```
