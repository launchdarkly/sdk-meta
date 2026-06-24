---
id: js-client-sdk/sdk-docs/features/flagchanges/flag-changes-all-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: All-flags change subscription example for JavaScript SDK v4.0.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
client.on('change', (context, changedKeys) => {
  changedKeys.forEach(flagKey => {
    const flagValue = client.variation(flagKey, defaultValue);
  });
});
```
