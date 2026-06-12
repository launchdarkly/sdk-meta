---
id: node-server-sdk/sdk-docs/features/aimetrics/delayed-feedback
sdk: node-server-sdk
kind: reference
lang: typescript
description: Capture tracker metadata at generation time and send delayed feedback with ldClient.track for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
// Generation time
const aiConfig = await aiClient.completionConfig(
  aiConfigKey,
  context,
  defaultValue,
  variables
);
const trackData = aiConfig.tracker.getTrackData();

// Persist trackData together with the context

// Feedback time
ldClient.track(
  feedback.kind === "positive"
    ? "$ld:ai:feedback:user:positive"
    : "$ld:ai:feedback:user:negative",
  context, // context from generation time
  persistedTrackData,
  1
);
```
