---
id: python-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: python-server-sdk
kind: reference
lang: python
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient
from ldclient.config import Config
from ldclient import datasystem

relay_uri = "http://my-relay-proxy:8030"

ldclient.set_config(
    Config(
        "YOUR_SDK_KEY",
        datasystem_config=datasystem.custom()
            .initializers([
                datasystem.polling_ds_builder().base_uri(relay_uri),
                datasystem.polling_ds_builder(),
            ])
            .synchronizers(
                datasystem.streaming_ds_builder().base_uri(relay_uri),
                datasystem.streaming_ds_builder(),
                datasystem.polling_ds_builder(),
            )
            .build(),
    )
)

client = ldclient.get()
```
