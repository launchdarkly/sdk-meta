---
id: fastly-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/fastly-syntax-only

---

```typescript
const flagValue = await client.variation('example-flag-key', context, false);
```
