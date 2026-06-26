---
id: python-server-sdk/sdk-docs/features/datasaving/file-bootstrap
sdk: python-server-sdk
kind: reference
lang: python
description: Data saving mode with file-based bootstrap and live updates for Python.
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
        datasystem_config=datasystem.custom()
            .initializers([
                datasystem.file_ds_builder(paths=["flags.json"]),
                datasystem.polling_ds_builder(),
            ])
            .synchronizers(
                datasystem.streaming_ds_builder(),
                datasystem.polling_ds_builder(),
            )
            .build(),
    )
)

client = ldclient.get()
```
