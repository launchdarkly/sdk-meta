---
id: react-client-sdk/sdk-docs/features/logging/logging
sdk: react-client-sdk
kind: reference
lang: javascript
description: basicLogger debug-level configuration example for React Web.
validation:
  scaffold: react-client-sdk/scaffolds/react-syntax-only

---

```javascript
import { basicLogger } from 'launchdarkly-js-client-sdk';

export default withLDProvider({
  clientSideID: 'example-client-side-id',
  options: {
    logger: basicLogger({level: 'debug'})
  }
})(App);
```
