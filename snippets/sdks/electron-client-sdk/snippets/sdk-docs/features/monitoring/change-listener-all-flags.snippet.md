---
id: electron-client-sdk/sdk-docs/features/monitoring/change-listener-all-flags
sdk: electron-client-sdk
kind: reference
lang: javascript
description: All-flags change listener for Electron (JavaScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
client.on('change', (allFlagChanges) => {
  Object.keys(allFlagChanges).forEach((key) => {
    console.log('Flag ' + key + ' is now ' + allFlagChanges[key]);
  });
});
```
