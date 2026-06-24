---
id: node-client-sdk/sdk-docs/features/flagchanges/flag-changes-specific
sdk: node-client-sdk
kind: reference
lang: javascript
description: Single-flag change subscription example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
client.on('change:example-flag-key', (context) => {
  const newValue = client.variation('example-flag-key', false);
  console.log('example-flag-key changed to:', newValue);
});
```
