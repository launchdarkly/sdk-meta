---
id: vercel-server-sdk/sdk-docs/features/events/flush
sdk: vercel-server-sdk
kind: reference
lang: typescript
description: Event flush inside waitUntil example for Vercel SDK v1.2.0+.
validation:
  scaffold: vercel-server-sdk/scaffolds/vercel-syntax-only

---

```typescript
import { waitUntil } from '@vercel/functions';

waitUntil(
  client.flush((err, res) => {
    console.log(`flushed events result: ${res}, error: ${err}`);
  }),
);
```
