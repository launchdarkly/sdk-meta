---
id: python-server-sdk/sdk-docs/features/testdata/flag-behavior-v8
sdk: python-server-sdk
kind: reference
lang: python
description: Configuring test data flag behavior for Python SDK v8.0.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
# This flag is true for the context with the key "example-context-key" and kind
# of "organization", and false for everyone else.

td.update(
    td.flag("flag-key-456def") \
        .variation_for_key("organization", "example-context-key", True) \
        .fallthrough_variation(False)
)

# This flag returns the string variation "green" for contexts of the kind "user"
# who have the custom attribute "admin" with a value of true, and "red" for
# everyone else.
td.update(
    td.flag("flag-key-789ghi") \
        .variations("red", "green")
        .fallthrough_variation(0)
        .if_match_context("user", "admin", True)
        .then_return(1)
)
```
