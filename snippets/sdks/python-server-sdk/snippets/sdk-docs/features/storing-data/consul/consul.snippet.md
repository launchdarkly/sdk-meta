---
id: python-server-sdk/sdk-docs/features/storing-data/consul/consul
sdk: python-server-sdk
kind: reference
lang: python
description: Consul feature store configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config
from ldclient.feature_store import CacheConfig
from ldclient.integrations import Consul

store = Consul.new_feature_store(host='my-consul', port=8100,
    prefix='my-key-prefix', caching=CacheConfig(expiration=30))

config = Config(feature_store=store)
ldclient.set_config(config)
```
