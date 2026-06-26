---
id: python-server-sdk/sdk-docs/features/agentcontrol/agent-config
sdk: python-server-sdk
kind: reference
lang: python
description: Customize an AgentControl config in agent mode for Python AI.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
from ldai import AIAgentConfigDefault

agent = aiclient.agent_config(
  'example-config-key',
  context,
  AIAgentConfigDefault(
    enabled=False
  ),
  { 'example_custom_variable': 'example_custom_value'}
)
```
