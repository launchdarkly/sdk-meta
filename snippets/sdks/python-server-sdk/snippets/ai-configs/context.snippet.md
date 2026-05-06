---
id: python-server-sdk/ai-configs/context
sdk: python-server-sdk
kind: context
lang: python
file: python-server-sdk/ai-configs/context.txt
description: Build an evaluation context for python-server-sdk AI Configs.
---

```python
context = Context.builder("context-key-123abc")
    .set("firstName", "Sandy")
    .set("lastName", "Smith")
    .set("email", "sandy@example.com")
    .set("groups", ["Google", "Microsoft"])
    .build()
```
