---
id: python-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: python-server-sdk
kind: reference
lang: python
description: Flag evaluation example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
show_feature = ldclient.get().variation("your.feature.key", context, False)
```
