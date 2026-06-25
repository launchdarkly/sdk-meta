---
id: electron-client-sdk/sdk-docs/features/bootstrapping/bootstrapping-js
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Bootstrapping example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```js
const client = LaunchDarkly.initializeInMain(
  'example-client-side-id',
  user,
  {
    bootstrap: {
      flagKey1: flagValue1,
      flagKey2: flagValue2
    }
  }
);
```
