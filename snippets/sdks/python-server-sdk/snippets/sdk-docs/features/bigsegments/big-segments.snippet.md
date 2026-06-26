---
id: python-server-sdk/sdk-docs/features/bigsegments/big-segments
sdk: python-server-sdk
kind: reference
lang: python
description: Big segments Redis store configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
import ldclient
from ldclient.config import Config, BigSegmentsConfig
from ldclient.integrations import Redis

store = Redis.new_big_segment_store(
    url='redis://your-redis:6379',
    prefix='example-client-side-id')

config = Config("YOUR_SDK_KEY", big_segments=BigSegmentsConfig(store=store))
ldclient.set_config(config)
client = ldclient.get()
```
