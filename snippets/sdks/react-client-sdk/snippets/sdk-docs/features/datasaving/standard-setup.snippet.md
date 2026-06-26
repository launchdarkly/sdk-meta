---
id: react-client-sdk/sdk-docs/features/datasaving/standard-setup
sdk: react-client-sdk
kind: reference
lang: javascript
description: Data saving mode standard setup for React Web.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

export const LDReactProvider = createLDReactProvider('example-client-side-id', context, {
  ldOptions: { dataSystem: {} },
});
```
