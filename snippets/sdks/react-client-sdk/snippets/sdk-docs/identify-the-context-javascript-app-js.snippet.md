---
id: react-client-sdk/sdk-docs/identify-the-context-javascript-app-js
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript: app.js in section \"Identify the context\""
---

```js
// app.js
import React, { useEffect } from 'react';
import { useFlags, useLDClient } from 'launchdarkly-react-client-sdk';

export default function App {
  const flags = useFlags();
  const ldClient = useLDClient();

  useEffect(() => {
    ldClient.identify({ key: 'example-context-key' });
  }, []);

  return <div>Let your feature flags fly!</div>
}
```
