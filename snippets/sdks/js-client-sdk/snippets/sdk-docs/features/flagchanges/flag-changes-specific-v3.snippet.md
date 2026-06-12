---
id: js-client-sdk/sdk-docs/features/flagchanges/flag-changes-specific-v3
sdk: js-client-sdk
kind: reference
lang: javascript
description: Single-flag change subscription example for JavaScript SDK v3.x.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
client.on('change:example-flag-key', (context) => {
  const newValue = client.variation('example-flag-key', false);
  console.log('example-flag-key changed to:', newValue);
});
```
