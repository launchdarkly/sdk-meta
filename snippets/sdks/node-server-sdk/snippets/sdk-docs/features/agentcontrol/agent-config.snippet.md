---
id: node-server-sdk/sdk-docs/features/agentcontrol/agent-config
sdk: node-server-sdk
kind: reference
lang: typescript
description: Customize an AgentControl config in agent mode for Node.js (server-side) AI.
validation:
  scaffold: node-server-sdk/scaffolds/node-syntax-only

---

```ts
const fallbackConfig = { enabled: false };

const agent: LDAIAgentConfig = await aiClient.agentConfig(
  'example-config-key',
  context,
  fallbackConfig,
  { 'exampleCustomVariable': 'exampleCustomValue' },
);
```
