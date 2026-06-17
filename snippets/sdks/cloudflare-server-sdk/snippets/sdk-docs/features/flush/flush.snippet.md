---
id: cloudflare-server-sdk/sdk-docs/features/flush/flush
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Manual event flush example for Cloudflare SDK v2.3.0+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
// executionContext is the Cloudflare worker handler context
// https://github.com/cloudflare/workers-types/blob/master/index.d.ts#L567
executionContext.waitUntil(
  client.flush((err, res) => {
    console.log(`flushed events result: ${res}, error: ${err}`);
  }),
);
```
