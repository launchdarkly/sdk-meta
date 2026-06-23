---
id: react-client-sdk/sdk-docs/features/identify/identify
sdk: react-client-sdk
kind: reference
lang: javascript
description: Identify example for the React SDK.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only
---

```javascript
import { useLDClient } from 'launchdarkly-react-client-sdk';

let ldClient = useLDClient();

ldClient.identify(newContext, null, () => {
  console.log("New context's flags available");
});
```
