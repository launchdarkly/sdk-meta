---
id: python-server-sdk/sdk-docs/migration-7-to-8-understanding-differences-between-users-and-contexts-8-0-syntax-multi-context
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```python
multi_context = Context.create_multi(
    Context.create("example-user-key"),
    Context.create("example-device-key", "device")
)
```
