---
id: node-client-sdk/sdk-docs/features/evaluating/evaluating-ts
sdk: node-client-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Node.js (client-side).
validation:
  scaffold: node-client-sdk/scaffolds/node-client-syntax-only

---

```typescript
const boolFlagValue = client.boolVariation('example-flag-key', false);
const numberFlagValue = client.numberVariation('example-flag-key', 2);
const stringFlagValue = client.stringVariation('example-flag-key', 'default');
```
