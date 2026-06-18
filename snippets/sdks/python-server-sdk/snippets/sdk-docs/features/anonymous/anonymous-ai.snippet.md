---
id: python-server-sdk/sdk-docs/features/anonymous/anonymous-ai
sdk: python-server-sdk
kind: reference
lang: python
description: Anonymous context example for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
context = Context.builder("example-context-key").anonymous(True).build()
```
