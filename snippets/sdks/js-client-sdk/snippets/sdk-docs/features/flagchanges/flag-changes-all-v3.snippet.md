---
id: js-client-sdk/sdk-docs/features/flagchanges/flag-changes-all-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: All-flags change subscription example for JavaScript SDK v3.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
client.on('change', (context, changedKeys) => {
  console.log('flags changed:', changedKeys);
});
```
