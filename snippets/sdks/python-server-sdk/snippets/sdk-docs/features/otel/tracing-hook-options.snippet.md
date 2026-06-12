---
id: python-server-sdk/sdk-docs/features/otel/tracing-hook-options
sdk: python-server-sdk
kind: reference
lang: python
description: OpenTelemetry tracing hook with span and value options for the Python SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient
from ldclient.config import Config

from ldotel.tracing import Hook, HookOptions

config = Config(sdk_key, hooks=[Hook(HookOptions(add_spans=True, include_value=True))])
ldclient.set_config(config=config)
client = ldclient.get()
```
