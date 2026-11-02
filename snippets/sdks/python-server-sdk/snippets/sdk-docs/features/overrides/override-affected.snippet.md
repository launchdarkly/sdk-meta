---
id: python-server-sdk/sdk-docs/features/overrides/override-affected
sdk: python-server-sdk
kind: reference
lang: python
description: Check whether a flag override affected an evaluation in Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
detail = client.variation_detail("my-boolean-flag-key", context, False)
if detail.reason.get("overrideAffected"):
    print("An override affected this evaluation")
```
