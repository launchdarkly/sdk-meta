---
id: cloudflare-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Flag evaluation reason example for Cloudflare.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
const { value, variationIndex, reason } = await client.variationDetail(flagKey, context, false);
```
