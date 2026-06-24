---
id: python-server-sdk/sdk-docs/features/flagchanges/flag-changes
sdk: python-server-sdk
kind: reference
lang: python
description: Flag change subscription example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
def flag_change_listener(flag_change):
    print(f"{flag_change.key} has changed")


listener = ldclient.get().flag_tracker.add_listener(flag_change_listener)
```
