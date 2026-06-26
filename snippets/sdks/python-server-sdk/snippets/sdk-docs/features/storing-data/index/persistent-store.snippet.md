---
id: python-server-sdk/sdk-docs/features/storing-data/index/persistent-store
sdk: python-server-sdk
kind: reference
lang: python
description: Persistent feature store configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config

store = SomeKindOfFeatureStore(store_options)
config = Config(feature_store=store)
ldclient.set_config(config)
```
