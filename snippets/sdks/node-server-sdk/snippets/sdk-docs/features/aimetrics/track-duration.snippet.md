---
id: node-server-sdk/sdk-docs/features/aimetrics/track-duration
sdk: node-server-sdk
kind: reference
lang: typescript
description: Track duration manually for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
// Track your own start and stop time.

// Set duration to the time (in ms) that your AI model generation takes.
// The duration may include network latency, depending on how you calculate it.

aiConfig.tracker.trackDuration(duration);
```
