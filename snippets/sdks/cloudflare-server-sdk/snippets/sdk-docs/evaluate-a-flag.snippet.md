---
id: cloudflare-server-sdk/sdk-docs/evaluate-a-flag
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
file: cloudflare-server-sdk/sdk-docs/evaluate-a-flag.ts
description: "Cloudflare edge SDK in section \"Evaluate a flag\""
validation:
  scaffold: cloudflare-server-sdk/scaffolds/edge-cloudflare-eval
---

```typescript
const context = {
   "kind": 'user',
   "key": 'example-user-key',
   "name": 'Sandy'
};

const flagValue = await client.variation('example-flag-key', context, false);
```
