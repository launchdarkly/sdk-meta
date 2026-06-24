---
id: js-client-sdk/sdk-docs/features/flagchanges/flag-changes-specific-v4
sdk: js-client-sdk
kind: reference
lang: javascript
description: Single-flag change subscription example for JavaScript SDK v4.0.
validation:
  scaffold: js-client-sdk/scaffolds/js-syntax-only

---

```javascript
client.on('change:example-flag-key', (context) => {
  const flagValue = client.variation('example-flag-key', false);
});
```
