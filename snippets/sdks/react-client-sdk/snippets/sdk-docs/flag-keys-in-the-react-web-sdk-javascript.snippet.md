---
id: react-client-sdk/sdk-docs/flag-keys-in-the-react-web-sdk-javascript
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Flag keys and the deprecated `useFlags` hook\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

const LDProvider = createLDReactProvider(
  'example-client-side-id',
  { kind: 'user', key: 'example-user-key' },
  {
    ldOptions: {
      useCamelCaseFlagKeys: false,
    },
  },
);
```
