---
id: python-server-sdk/sdk-docs/features/contextconfig/context-example-ai
sdk: python-server-sdk
kind: reference
lang: python
description: Context example for Python AI SDK.
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
