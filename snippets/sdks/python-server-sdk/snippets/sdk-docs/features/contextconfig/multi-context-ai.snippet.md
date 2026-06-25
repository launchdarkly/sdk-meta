---
id: python-server-sdk/sdk-docs/features/contextconfig/multi-context-ai
sdk: python-server-sdk
kind: reference
lang: python
description: Multi-context example for Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
multi_context = Context.create_multi(
    Context.create("example-user-key"),
    Context.create("example-device-key", "device")
)
```
