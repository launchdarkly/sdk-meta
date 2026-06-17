---
id: node-client-sdk/sdk-docs/features/flush/flush
sdk: node-client-sdk
kind: reference
lang: javascript
description: Manual event flush example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```javascript
client.flush();

// or, with a callback:
client.flush(() => {
  console.log('flush complete');
});

// or, with a Promise:
client.flush().then(() => {
  console.log('flush complete');
});
```
