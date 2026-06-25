---
id: electron-client-sdk/sdk-docs/features/localstorage/localstorage-js
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Local storage caching example for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```js
const client = LaunchDarkly.initializeInMain(
  'example-client-side-id',
  user,
  {
    bootstrap: 'localStorage'
  }
);
```
