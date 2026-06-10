---
id: node-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-ts
sdk: node-server-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Node.js (server-side, TypeScript).
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
client.variationDetail('example-flag-key', context, false,
  (detail) => {
    const detailValue = detail.value;
    const detailIndex = detail.variationIndex;
    const detailReason = detail.reason;
});
```
