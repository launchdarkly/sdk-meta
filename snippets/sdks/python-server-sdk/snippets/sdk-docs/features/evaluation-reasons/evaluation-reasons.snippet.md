---
id: python-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: python-server-sdk
kind: reference
lang: python
description: Flag evaluation reason example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
detail = client.variation_detail("example-flag-key", my_context, False)
value = detail.value
index = detail.variation_index
reason = detail.reason
```
