---
id: node-server-sdk/sdk-docs/features/aimetrics/bedrock-completion
sdk: node-server-sdk
kind: reference
lang: typescript
description: Record metrics from a Bedrock Converse command in completion mode for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```typescript
const { tracker } = aiConfig;

if (aiConfig.enabled) {

  // Pass in the result of the Bedrock Converse command.
  // When you call the Bedrock Converse command, use details from aiConfig.
  // For instance, you can pass aiConfig.messages
  // and aiConfig.model to your specific Bedrock Converse command.

  const completion = tracker.trackBedrockConverseMetrics(
    await awsClient.send(
      new ConverseCommand({
        modelId: aiConfig.model?.name ?? 'no-model',
        messages: mapPromptToConversation(aiConfig.messages ?? []),
        inferenceConfig: {
          temperature: (aiConfig.model?.parameters?.temperature as number) ?? 0.5,
          maxTokens: (aiConfig.model?.parameters?.maxTokens as number) ?? 4096,
        },
      }),
    ),
  );

} else {

  // Application path to take when the aiConfig is disabled

}
```
