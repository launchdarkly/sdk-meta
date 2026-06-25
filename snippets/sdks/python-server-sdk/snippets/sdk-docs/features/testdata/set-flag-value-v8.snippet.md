---
id: python-server-sdk/sdk-docs/features/testdata/set-flag-value-v8
sdk: python-server-sdk
kind: reference
lang: python
description: Setting a test data flag to a specific value for Python SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
td.update(td.flag("example-flag-key").variation_for_all(True))
```
