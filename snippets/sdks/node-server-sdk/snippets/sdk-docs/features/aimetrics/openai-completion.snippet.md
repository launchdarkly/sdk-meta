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
const { tracker } = aiConfig;

if (aiConfig.enabled) {

  // Pass in the result of the OpenAI operation.
  // When you call the OpenAI operation, use details from aiConfig.
  // For instance, you can pass aiConfig.messages
  // and aiConfig.model to your specific OpenAI operation.
  //
  // CAUTION: If the call inside of trackOpenAIMetrics throws an exception,
  // the SDK will re-throw that exception

  const completion = await tracker.trackOpenAIMetrics(async () =>
    client.chat.completions.create({
      messages: aiConfig.messages || [],
      model: aiConfig.model?.name || 'gpt-4',
      temperature: (aiConfig.model?.parameters?.temperature as number) ?? 0.5,
      maxTokens: (aiConfig.model?.parameters?.maxTokens as number) ?? 4096,
    }),
  );

} else {

  // Application path to take when the aiConfig is disabled

}
```
