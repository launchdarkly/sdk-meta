---
id: python-server-sdk/sdk-docs/features/hooks/add-hook
sdk: python-server-sdk
kind: reference
lang: python
description: Adding a hook to an existing client for the Python SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
ldclient.set_config(config=config)
client = ldclient.get()

client.add_hook(example_hook)
```
