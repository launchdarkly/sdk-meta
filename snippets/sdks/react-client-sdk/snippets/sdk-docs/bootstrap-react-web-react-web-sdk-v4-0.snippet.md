---
id: react-client-sdk/sdk-docs/bootstrap-react-web-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"React Web\""
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

// `flags` is the result of your server-side SDK call to get all flags,
// e.g. JSON.parse(bootstrapData) where bootstrapData is the
// server-rendered string. Pass it as a plain key-value object.
const LDProvider = createLDReactProvider(
  'example-client-side-id',
  { kind: 'user', key: 'example-user-key' },
  {
    bootstrap: flags,
  },
);
```
