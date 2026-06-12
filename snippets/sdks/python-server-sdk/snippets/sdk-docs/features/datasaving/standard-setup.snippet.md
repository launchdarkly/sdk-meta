---
id: python-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: python-server-sdk
kind: reference
lang: python
description: Data saving mode standard setup for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient
from ldclient.config import Config
from ldclient import datasystem

ldclient.set_config(
    Config(
        "YOUR_SDK_KEY",
        datasystem_config=datasystem.default().build(),
    )
)

client = ldclient.get()
```
