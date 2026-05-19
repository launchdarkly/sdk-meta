---
id: python-server-sdk/ai-configs/implementation
sdk: python-server-sdk
kind: implementation
lang: python
file: python-server-sdk/ai-configs/implementation.txt
description: Resolve an AI Config with a fallback for python-server-sdk.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
fallback_value = AIConfig(
  enabled=True,
  model=ModelConfig(
      name="my-default-model",
      parameters={"temperature": 0.8},
  ),
  messages=[LDMessage(role="system", content="")],
  provider=ProviderConfig(name="my-default-provider"),
)

config, tracker = aiclient.config('{{configKey}}', context, fallback_value, { 'example_custom_variable': 'example_custom_value'})
```
