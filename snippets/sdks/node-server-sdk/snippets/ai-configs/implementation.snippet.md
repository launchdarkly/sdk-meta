---
id: node-server-sdk/ai-configs/implementation
sdk: node-server-sdk
kind: implementation
lang: javascript
file: node-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for node-server-sdk.
---

```javascript
const fallbackConfig = {
  model: {
    name: 'my-default-model',
    parameters: { temperature: 0.8 }
  },
  messages: [ { role: 'system', content: '' } ],
  provider: { name: 'my-default-provider' },
  enabled: true,
};

const aiConfig: LDAIConfig = aiClient.config(
  '{{configKey}}',
  context,
  fallbackConfig,
  { 'exampleCustomVariable': 'exampleCustomValue' },
);
```
