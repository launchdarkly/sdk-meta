---
id: vercel-server-sdk/sdk-docs/features/flush/flush
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Manual event flush example for Vercel SDK v1.2.0+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
context.waitUntil(
  client.flush((err, res) => {
    console.log(`flushed events result: ${res}, error: ${err}`);
  }),
);
```
