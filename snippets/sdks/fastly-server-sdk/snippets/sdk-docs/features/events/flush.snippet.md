---
id: fastly-server-sdk/sdk-docs/features/events/flush
sdk: fastly-server-sdk
kind: reference
lang: typescript
description: Event flush inside waitUntil example for Fastly.
validation:
  scaffold: fastly-server-sdk/scaffolds/fastly-syntax-only

---

```typescript
async function handleRequest(event: FetchEvent) {
  // ...

  event.waitUntil(
    client.flush((err, res) => {
      console.log(`flushed events result: ${res}, error: ${err}`);
    }),
  );

  // ...
}
```
