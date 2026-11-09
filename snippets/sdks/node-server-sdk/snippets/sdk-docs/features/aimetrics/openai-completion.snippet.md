---
id: node-server-sdk/sdk-docs/features/aimetrics/openai-completion
sdk: node-server-sdk
kind: reference
lang: typescript
description: Record metrics from an OpenAI operation in completion mode for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
import { getAIMetricsFromResponse } from '@launchdarkly/server-sdk-ai-openai';

// Create a new tracker for each AI run
const tracker = aiConfig.createTracker();

if (aiConfig.enabled) {

  // Pass in a function that calls the OpenAI operation.
  // When you call the OpenAI operation, use details from aiConfig.
  // For instance, you can pass aiConfig.messages
  // and aiConfig.model to your specific OpenAI operation.
  //
  // CAUTION: If the call inside of trackMetricsOf throws an exception,
  // the SDK records an error and re-throws that exception.

  const completion = await tracker.trackMetricsOf(getAIMetricsFromResponse, async () =>
    client.chat.completions.create({
      messages: aiConfig.messages || [],
      model: aiConfig.model?.name || 'gpt-4',
      temperature: (aiConfig.model?.parameters?.temperature as number) ?? 0.5,
      max_tokens: (aiConfig.model?.parameters?.maxTokens as number) ?? 4096,
    }),
  );

} else {

  // Application path to take when the aiConfig is disabled

}
```
