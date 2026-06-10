---
id: akamai-server-edgekv-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: akamai-server-edgekv-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Akamai.
validation:
  scaffold: akamai-server-edgekv-sdk/scaffolds/akamai-syntax-only

---

```typescript
const { value, variationIndex, reason } = await client.variationDetail(flagKey, context, false);
```
