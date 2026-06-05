---
id: python-server-sdk/sdk-docs/features/config/app-config
sdk: python-server-sdk
kind: reference
lang: python
description: Application metadata configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
config = Config(sdk_key='YOUR_SDK_KEY',
  application = {"id": "authentication-service", "version": "1.0.0"})
ldclient.set_config(config)

```
