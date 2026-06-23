---
id: python-server-sdk/sdk-docs/features/logging/logging
sdk: python-server-sdk
kind: reference
lang: python
description: Debug log level configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
ld_logger = logging.getLogger("ldclient")
ld_logger.setLevel(logging.DEBUG)
```
