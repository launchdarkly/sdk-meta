---
id: python-server-sdk/sdk-docs/features/anonymous/anonymous-v8
sdk: python-server-sdk
kind: reference
lang: python
description: Anonymous context example for Python, SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
context = Context.builder("example-context-key").anonymous(True).build()
```
