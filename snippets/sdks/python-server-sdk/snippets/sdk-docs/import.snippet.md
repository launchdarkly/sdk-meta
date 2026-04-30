---
id: python-server-sdk/sdk-docs/import
sdk: python-server-sdk
kind: reference
lang: python
description: SDK + observability plugin imports.
validation:
  scaffold: python-server-sdk/scaffolds/with-test-data
---

```python
import ldclient
from ldclient.config import Config

# optional observability plugin, requires Python SDK v9.12+
import ldobserve
from ldobserve import ObservabilityConfig, ObservabilityPlugin, observe
```
