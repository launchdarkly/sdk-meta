---
id: node-server-sdk/sdk-docs/features/aimetrics/track-feedback
sdk: node-server-sdk
kind: reference
lang: typescript
description: Track output satisfaction rate for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import { LDFeedbackKind } from '@launchdarkly/server-sdk-ai';

// Track your own output satisfaction rate.

// Pass in LDFeedbackKind.Positive or LDFeedbackKind.Negative.
aiConfig.tracker.trackFeedback({ kind: LDFeedbackKind.Positive });

```
