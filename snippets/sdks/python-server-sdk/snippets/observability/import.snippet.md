---
id: python-server-sdk/observability/import
sdk: python-server-sdk
kind: import
lang: python
file: python-server-sdk/observability/import.txt
description: Import statements for python-server-sdk observability plugin.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import logging

import ldclient
from ldclient.config import Config
from ldobserve import ObservabilityConfig, ObservabilityPlugin, observe
```
