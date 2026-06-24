---
id: node-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: node-server-sdk
kind: reference
lang: javascript
description: Flag update event subscription example for Node.js (server-side).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```javascript
client.on('update', (param) => {
  console.log('a flag was changed: ' + param.key);
});

client.on('update:example-flag-key', () => {
  console.log('the example-flag-key flag was changed');
});
```
