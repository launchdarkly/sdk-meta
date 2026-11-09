---
id: node-server-sdk/sdk-docs/features/aimetrics/vercel-generate-text
sdk: node-server-sdk
kind: reference
lang: typescript
description: Record metrics from a Vercel AI SDK generateText operation for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import { generateText } from 'ai';
import {
  VercelRunnerFactory,
  getAIMetricsFromResponse,
} from '@launchdarkly/server-sdk-ai-vercel';

const { tracker } = aiConfig;

// Create a Vercel AI SDK model from the provider and model in aiConfig.
// Install the Vercel AI SDK provider package, such as @ai-sdk/openai,
// for each provider that your config uses.
//
// CAUTION: createVercelModel throws an exception if the provider
// is not supported or its Vercel AI SDK package is not installed.
const model = await VercelRunnerFactory.createVercelModel(aiConfig);

// Pass in a function that calls the Vercel AI SDK's generateText function.
// trackMetricsOf records the duration, token usage, and success or error.
const completion = await tracker.trackMetricsOf(getAIMetricsFromResponse, () =>
  generateText({
    model,
    messages: aiConfig.messages ?? [],
    ...VercelRunnerFactory.mapParameters(aiConfig.model?.parameters),
  }),
);
```
