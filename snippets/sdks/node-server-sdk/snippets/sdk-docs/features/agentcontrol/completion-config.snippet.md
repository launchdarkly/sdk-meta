---
id: node-server-sdk/sdk-docs/features/agentcontrol/completion-config
sdk: node-server-sdk
kind: reference
lang: typescript
description: Customize an AgentControl config in completion mode for Node.js (server-side) AI.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
const fallbackConfig = { enabled: false };

const aiConfig: LDAIConfig = await aiClient.completionConfig(
  'example-config-key',
  context,
  fallbackConfig,
  { exampleCustomVariable: 'exampleCustomValue' },
);

if (aiConfig.enabled) {
  // Send a request to your AI provider using the customized config
} else {
  // Application path to take when the config is disabled
}
```
