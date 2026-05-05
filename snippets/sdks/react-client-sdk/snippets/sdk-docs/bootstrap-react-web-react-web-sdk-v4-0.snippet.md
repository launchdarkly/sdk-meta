---
id: react-client-sdk/sdk-docs/bootstrap-react-web-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"React Web\""
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
