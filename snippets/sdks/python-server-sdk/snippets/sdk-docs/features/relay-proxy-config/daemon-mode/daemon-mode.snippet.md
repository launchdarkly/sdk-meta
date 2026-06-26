---
id: python-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: python-server-sdk
kind: reference
lang: python
description: Daemon mode configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
store = SomeKindOfFeatureStore(store_options)

config = Config(
  feature_store=store,
  use_ldd=True)
```
