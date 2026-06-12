---
id: python-server-sdk/sdk-docs/features/storing-data/dynamodb/dynamodb
sdk: python-server-sdk
kind: reference
lang: python
description: DynamoDB feature store configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config
from ldclient.feature_store import CacheConfig
from ldclient.integrations import DynamoDB

store = DynamoDB.new_feature_store('my_table',
    caching=CacheConfig(expiration=30))

config = Config(feature_store=store)
ldclient.set_config(config)
```
