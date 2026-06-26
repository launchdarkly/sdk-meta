---
id: node-server-sdk/sdk-docs/features/aimetrics/vercel-stream-text
sdk: node-server-sdk
kind: reference
lang: typescript
description: Record metrics from a Vercel AI SDK streamText operation for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
const { tracker } = aiConfig;

// Pass in the result of the Vercel AI SDK's streamText function.
// When you call streamText, use details from the aiConfig,
// mapped to the input format required for the Vercel AI SDK.
//
// CAUTION: The toVercelAISDK function may throw an exception
// if a Vercel AI SDK model cannot be determined.

const completion = tracker.trackVercelAISDKStreamTextMetrics(() =>
 streamText(
    aiConfig.toVercelAISDK(vercelProvider, vercelProviderOptions)
  )
)

```
