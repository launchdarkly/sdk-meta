---
id: python-server-sdk/sdk-docs/features/flagchanges/flag-value-changes
sdk: python-server-sdk
kind: reference
lang: python
description: Flag value change subscription example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
def flag_value_change_listener(flag_change):
    print(f"{flag_change.key} has changed from {flag_change.old_value} to {flag_change.new_value}")


listener = ldclient.get().flag_tracker.add_flag_value_change_listener('example-flag-key', context, flag_value_change_listener)
```
