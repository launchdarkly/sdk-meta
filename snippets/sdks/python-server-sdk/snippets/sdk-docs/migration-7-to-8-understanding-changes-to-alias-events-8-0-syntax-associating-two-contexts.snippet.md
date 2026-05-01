---
id: python-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-alias-events-8-0-syntax-associating-two-contexts
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, associating two contexts in section \"Understanding changes to alias events\""
---

```python
context1 = Context.create("example-user-key")
context2 = Context.create("example-device-key", "device")
multi_context = Context.create_multi(context1, context2)
client.identify(multi_context)
```
