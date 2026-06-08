---
id: akamai-server-edgekv-sdk/sdk-docs/features/evaluating/evaluating
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Akamai.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only

---

```typescript
const flagValue = await client.variation('example-flag-key', context, false);
```
