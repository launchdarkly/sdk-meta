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
import { streamText } from 'ai';
import {
  VercelRunnerFactory,
  getAIMetricsFromStream,
} from '@launchdarkly/server-sdk-ai-vercel';

const { tracker } = aiConfig;

// Create a Vercel AI SDK model from the provider and model in aiConfig.
// Install the Vercel AI SDK provider package, such as @ai-sdk/openai,
// for each provider that your config uses.
//
// CAUTION: createVercelModel throws an exception if the provider
// is not supported or its Vercel AI SDK package is not installed.
const model = await VercelRunnerFactory.createVercelModel(aiConfig);

// Pass in a function that calls the Vercel AI SDK's streamText function.
// trackStreamMetricsOf returns the stream immediately, and records
// the duration, token usage, and success or error when the stream finishes.
const stream = tracker.trackStreamMetricsOf(
  () =>
    streamText({
      model,
      messages: aiConfig.messages ?? [],
      ...VercelRunnerFactory.mapParameters(aiConfig.model?.parameters),
    }),
  getAIMetricsFromStream,
);
```
