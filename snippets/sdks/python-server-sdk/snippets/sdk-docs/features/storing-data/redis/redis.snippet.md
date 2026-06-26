---
id: python-server-sdk/sdk-docs/features/storing-data/redis/redis
sdk: python-server-sdk
kind: reference
lang: python
description: Redis feature store configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config
from ldclient.feature_store import CacheConfig
from ldclient.integrations import Redis

store = Redis.new_feature_store(url='redis://my-redis:6379',
    prefix='my-key-prefix', caching=CacheConfig(expiration=30))

config = Config(feature_store=store)
ldclient.set_config(config)
```
