---
id: python-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: python-server-sdk
kind: reference
lang: python
description: Offline mode example for Python SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# Initialization:
ldclient.set_config(Config("YOUR_SDK_KEY", offline=True))
ldclient.get().variation("any.feature.flag", context, False) # will always return the default value (false)

```
