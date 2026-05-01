---
id: python-server-sdk/sdk-docs/migration-7-to-8-working-with-built-in-and-custom-attributes-8-0-syntax-context-with-attributes
sdk: python-server-sdk
kind: reference
lang: python
description: "8.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```python
# use Context.builder to specify attributes besides the key and kind

context = Context.builder("example-user-key") \
    .name("Sandy") \
    .set("email", "sandy@example.com") \
    .set("groups", ["admin"]) \
    .build()

# alternatively, use Context.from_dict to create a Context
# from properties in a dictionary, corresponding to the JSON representation of a context

pre_existing_dict = {
    'key': 'user-key-456def',
    'kind': 'user',
    'name': 'Sandy',
    'email': 'sandy@example.com',
    'groups': ['admin'],
}

context2 = Context.from_dict(pre_existing_dict)

```
