---
id: python-server-sdk/sdk-docs/migration-7-to-8-referencing-properties-of-an-attribute-object-8-0-syntax-context-with-object-attributes
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```python
address_data = {"street": "Main St", "city": "Springfield"}
context = Context.builder("example-user-key") \
    .set("address", address_data) \
    .build()
```
