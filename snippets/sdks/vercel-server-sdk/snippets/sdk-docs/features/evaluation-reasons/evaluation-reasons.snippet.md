---
id: vercel-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Vercel.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
const { value, variationIndex, reason } = await client.variationDetail(flagKey, context, false);
```
