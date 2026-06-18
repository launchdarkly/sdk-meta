---
id: python-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: python-server-sdk
kind: reference
lang: python
description: Context example for Python SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
context = Context.builder("example-context-key") \
    .set("firstName", "Sandy") \
    .set("lastName", "Smith") \
    .set("email", "sandy@example.com") \
    .set("groups", ["Acme", "Global Health Services"]) \
    .build()
```
