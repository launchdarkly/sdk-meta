---
id: node-server-sdk/sdk-docs/features/aimetrics/bedrock-completion
sdk: node-server-sdk
kind: reference
lang: typescript
description: Record metrics from a Bedrock Converse command in completion mode for the Node.js (server-side) AI SDK.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only-typescript

---

```typescript
import type { ConverseCommandOutput } from '@aws-sdk/client-bedrock-runtime';
import type { LDAIMetrics } from '@launchdarkly/server-sdk-ai';

// The SDK does not provide a metrics function for Bedrock.
// This function reads the success status and token usage
// from the result of the Bedrock Converse command.
function extractBedrockConverseMetrics(response: ConverseCommandOutput): LDAIMetrics {
  return {
    success: response.$metadata?.httpStatusCode === 200,
    tokens: response.usage
      ? {
          total: response.usage.totalTokens ?? 0,
          input: response.usage.inputTokens ?? 0,
          output: response.usage.outputTokens ?? 0,
        }
      : undefined,
  };
}

// Create a new tracker for each AI run
const tracker = aiConfig.createTracker();

if (aiConfig.enabled) {

  // Pass in a function that calls the Bedrock Converse command.
  // When you call the Bedrock Converse command, use details from aiConfig.
  // For instance, you can pass aiConfig.messages
  // and aiConfig.model to your specific Bedrock Converse command.
  //
  // CAUTION: If the call inside of trackMetricsOf throws an exception,
  // the SDK records an error and re-throws that exception.

  const completion = await tracker.trackMetricsOf(extractBedrockConverseMetrics, async () =>
    awsClient.send(
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
