---
id: python-server-sdk/sdk-docs/features/privateattrs/context-ai
sdk: python-server-sdk
kind: reference
lang: python
description: Marking context attributes private with the context builder for the Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
context = Context.builder("example-context-key") \
    .set("email", "sandy@example.com") \
    .private("email") \
    .build()
```
