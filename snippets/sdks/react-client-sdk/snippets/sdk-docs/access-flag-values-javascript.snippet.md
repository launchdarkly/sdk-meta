---
id: react-client-sdk/sdk-docs/access-flag-values-javascript
sdk: react-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Access flag values\""
---

```js
import { useBoolVariation, useLDClient } from '@launchdarkly/react-sdk';

function Home() {
  const ldClient = useLDClient();

  // You can call any of the methods from the JavaScript SDK
  // ldClient.identify({...})

  const devTestFlag = useBoolVariation('dev-test-flag', false);

  return devTestFlag ? <div>Flag on</div> : <div>Flag off</div>;
}

export default Home;
```
