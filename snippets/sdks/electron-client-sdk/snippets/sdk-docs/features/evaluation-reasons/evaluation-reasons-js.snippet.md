---
id: electron-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-js
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Flag evaluation reason example for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
const { value, variationIndex, reason } = client.variationDetail('example-flag-key', false);
```
