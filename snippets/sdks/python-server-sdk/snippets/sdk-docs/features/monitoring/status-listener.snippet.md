---
id: python-server-sdk/sdk-docs/features/monitoring/status-listener
sdk: python-server-sdk
kind: reference
lang: python
description: Data source status change listener for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
listener = ldclient.get().data_source_status_provider.add_listener(source_status_listener)
```
