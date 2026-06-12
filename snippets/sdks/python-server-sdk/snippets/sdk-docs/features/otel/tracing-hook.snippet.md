---
id: python-server-sdk/sdk-docs/features/otel/tracing-hook
sdk: python-server-sdk
kind: reference
lang: python
description: OpenTelemetry tracing hook configuration for the Python SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient
from ldclient import Config
from ldotel.tracing import Hook, HookOptions

config = Config(sdk_key, hooks=[Hook()])
ldclient.set_config(config=config)
client = ldclient.get()
```
