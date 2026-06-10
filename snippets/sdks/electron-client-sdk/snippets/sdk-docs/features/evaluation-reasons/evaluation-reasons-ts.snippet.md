---
id: electron-client-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons-ts
sdk: electron-client-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Electron (TypeScript).
validation:
  scaffold: electron-client-sdk/scaffolds/electron-syntax-only

---

```typescript
const { value, variationIndex, reason } = client.variationDetail('example-flag-key', false);
```
