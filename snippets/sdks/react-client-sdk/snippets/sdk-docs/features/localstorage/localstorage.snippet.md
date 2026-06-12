---
id: react-client-sdk/sdk-docs/features/localstorage/localstorage
sdk: react-client-sdk
kind: reference
lang: typescript
description: Local storage caching example for React Web.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```ts
import { createLDReactProvider } from '@launchdarkly/react-sdk';

const context = { kind: 'user', key: 'example-user-key' };

// Local storage is enabled by default
// You can optionally configure the maximum number of cached contexts (default is 5)
const LDProvider = createLDReactProvider(
  'example-client-side-id',
  context,
  {
    ldOptions: {
      maxCachedContexts: 3,
    },
  }
);
```
