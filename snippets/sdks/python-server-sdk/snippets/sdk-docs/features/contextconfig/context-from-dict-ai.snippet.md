---
id: python-server-sdk/sdk-docs/features/contextconfig/context-from-dict-ai
sdk: python-server-sdk
kind: reference
lang: python
description: Creating a context from a dictionary for Python AI SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
pre_existing_dict = {
    'key': 'example-context-key',
    'kind': 'user',
    'firstName': 'Sandy',
    'lastName': 'Smith',
    'email': 'sandy@example.com',
    'groups': ['Acme', 'Global Health Services'],
}

context = Context.from_dict(pre_existing_dict)
```
