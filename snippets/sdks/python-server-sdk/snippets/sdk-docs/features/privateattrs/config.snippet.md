---
id: python-server-sdk/sdk-docs/features/privateattrs/config
sdk: python-server-sdk
kind: reference
lang: python
description: Private attribute configuration for Python SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
# All attributes marked private
config = Config(all_attributes_private=True)

# Two attributes marked private
config = Config(private_attributes=["name", "email"])
```
