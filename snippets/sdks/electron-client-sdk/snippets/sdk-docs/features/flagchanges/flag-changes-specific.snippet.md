---
id: electron-client-sdk/sdk-docs/features/flagchanges/flag-changes-specific
sdk: electron-client-sdk
kind: reference
lang: javascript
description: Single-flag change subscription example for Electron.
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```javascript
client.on('change:example-flag-key', (context) => {
  const newValue = client.variation('example-flag-key', false);
  console.log('example-flag-key changed to: ' + newValue);
});
```
