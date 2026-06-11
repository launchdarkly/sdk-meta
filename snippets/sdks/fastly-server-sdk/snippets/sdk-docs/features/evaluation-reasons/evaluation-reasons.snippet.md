---
id: fastly-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/fastly-syntax-only

---

```typescript
const { value, variationIndex, reason } = await client.variationDetail('example-flag-key', context, false);
```
