---
id: python-server-sdk/sdk-docs/migration-7-to-8-understanding-differences-between-users-and-contexts-8-0-syntax-context-with-key
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```python
from ldclient import Context

# use Context.create when including only the key and kind
context = Context.create("example-context-key")
```
