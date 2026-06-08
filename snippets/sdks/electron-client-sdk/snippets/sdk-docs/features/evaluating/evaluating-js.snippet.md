---
id: electron-client-sdk/sdk-docs/features/evaluating/evaluating-js
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Flag evaluation example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
const flagValue = client.variation('example-flag-key', false);

// proceed based on flag value, for example:

if (flagValue)  {
  // feature flag targeting is on
} else {
  // feature flag targeting is off
}
```
