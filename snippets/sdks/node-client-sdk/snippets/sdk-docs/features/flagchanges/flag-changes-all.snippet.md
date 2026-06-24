---
id: node-client-sdk/sdk-docs/features/flagchanges/flag-changes-all
sdk: node-client-sdk
kind: reference
lang: javascript
description: All-flags change subscription example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
client.on('change', (context, changedKeys) => {
  console.log('flags changed:', changedKeys);
});
```
