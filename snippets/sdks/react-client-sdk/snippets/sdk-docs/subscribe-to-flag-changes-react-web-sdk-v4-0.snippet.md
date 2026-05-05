---
id: react-client-sdk/sdk-docs/subscribe-to-flag-changes-react-web-sdk-v4-0
sdk: react-client-sdk
kind: reference
lang: javascript
description: "React Web SDK v4.0 in section \"Subscribe to flag changes\""
---

```js
import { createLDReactProvider } from '@launchdarkly/react-sdk';

const LDProvider = createLDReactProvider(
  'example-client-side-id',
  { kind: 'user', key: 'example-user-key' },
  {
    ldOptions: {
      streaming: false, // or `true` to enable
    },
  },
);
```
