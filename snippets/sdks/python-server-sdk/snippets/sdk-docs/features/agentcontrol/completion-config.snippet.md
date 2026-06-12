---
id: python-server-sdk/sdk-docs/features/agentcontrol/completion-config
sdk: python-server-sdk
kind: reference
lang: python
description: Customize an AgentControl config in completion mode for Python AI.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
key = 'example-config-key'
context = Context.builder('example-context-key') \
  .kind('user') \
  .set('name', 'Sandy') \
  .build()
fallback_value = AICompletionConfigDefault(enabled=False)
variables = { 'example_custom_variable': 'example_custom_value' }

config = aiclient.completion_config(key, context, fallback_value, variables)

if config.enabled:
  # Send a request to your AI provider using the customized config
  pass
else:
  # Application path to take when the config is disabled
  pass
```
