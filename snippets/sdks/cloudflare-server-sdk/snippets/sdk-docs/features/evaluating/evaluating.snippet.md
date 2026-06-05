---
id: cloudflare-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Flag evaluation example for Cloudflare.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
const flagValue = await client.variation('example-flag-key', context, false);
```
