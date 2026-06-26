---
id: node-server-sdk/sdk-docs/features/aimetrics/track-time-to-first-token
sdk: node-server-sdk
kind: reference
lang: typescript
description: Track time to first token for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
// Track the time it takes to generate the first token

// Pass in the time (in ms) until your first token is generated
// This may include network latency, depending on how you calculate it

aiConfig.tracker.trackTimeToFirstToken(1000);
```
