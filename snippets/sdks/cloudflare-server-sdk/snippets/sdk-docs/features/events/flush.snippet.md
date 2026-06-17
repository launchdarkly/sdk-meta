---
id: cloudflare-server-sdk/sdk-docs/features/events/flush
sdk: cloudflare-server-sdk
kind: reference
lang: typescript
description: Event flush example for Cloudflare SDK v2.3.0+.
validation:
  scaffold: cloudflare-server-sdk/scaffolds/cloudflare-syntax-only

---

```typescript
executionContext.waitUntil(
  client.flush((err, res) => {
    console.log(`flushed events result: ${res}, error: ${err}`);
  }),
);
```
