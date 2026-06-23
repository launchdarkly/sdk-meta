---
id: electron-client-sdk/sdk-docs/features/monitoring/change-listener-single-flag
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Per-flag change listener for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
client.on('change:example-flag-key', (newValue, oldValue) => {
  console.log('The flag was ' + oldValue + ' and now it is ' + newValue);
});
```
