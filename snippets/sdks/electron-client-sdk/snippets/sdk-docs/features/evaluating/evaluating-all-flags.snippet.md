---
id: electron-client-sdk/sdk-docs/features/evaluating/evaluating-all-flags
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Fetching all flags example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
const flags = client.allFlags();
const flagValue = flags['example-flag-key'];
```
