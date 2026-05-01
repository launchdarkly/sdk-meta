---
id: python-server-sdk/sdk-docs/migration-7-to-8-understanding-changes-to-private-attributes-8-0-syntax-attribute-marked-private-for-one-context
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```python
context = Context.builder("example-context-key") \
    .name("Sandy") \
    .set("email", "sandy@example.com") \
    .private("email") \
    .build()
```
