---
id: react-client-sdk/sdk-docs/bootstrap-react-web-startoptions-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"React Web\" (via `startOptions`)"
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
