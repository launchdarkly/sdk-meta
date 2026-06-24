---
id: electron-client-sdk/sdk-docs/features/flagchanges/flag-changes-all
sdk: electron-client-sdk
kind: reference
lang: javascript
description: All-flags change subscription example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
client.on('change', (context, changedKeys) => {
  changedKeys.forEach((key) => {
    console.log('Flag ' + key + ' changed');
  });
});
```
