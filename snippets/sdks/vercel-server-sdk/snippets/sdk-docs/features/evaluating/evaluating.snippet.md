---
id: vercel-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Vercel.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
const flagValue = await client.variation('example-flag-key', context, false);
```
