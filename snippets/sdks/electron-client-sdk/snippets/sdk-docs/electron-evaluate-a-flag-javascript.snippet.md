---
id: electron-client-sdk/sdk-docs/electron-evaluate-a-flag-javascript
sdk: electron-client-sdk
kind: reference
lang: javascript
description: "JavaScript in section \"Evaluate a flag\""
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only
---

```js
const flagValue = client.variation('example-flag-key', false);

// proceed based on flag value, for example:

if (flagValue)  {
  // feature flag targeting is on
} else {
  // feature flag targeting is off
}
```
