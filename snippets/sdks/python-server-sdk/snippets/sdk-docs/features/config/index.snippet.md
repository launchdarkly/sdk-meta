---
id: python-server-sdk/sdk-docs/features/config/index
sdk: python-server-sdk
kind: reference
lang: python
description: SDK configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
config = Config(sdk_key='YOUR_SDK_KEY', http=HTTPConfig(connect_timeout=5))
ldclient.set_config(config)
```
