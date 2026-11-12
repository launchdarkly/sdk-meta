---
id: python-server-sdk/sdk-docs/features/overrides/file-overrides
sdk: python-server-sdk
kind: reference
lang: python
description: File-based flag override source configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient
from ldclient.config import Config
from ldclient import datasystem
from ldclient.integrations.overrides import FileOverrideSourceBuilder

overrides = FileOverrideSourceBuilder(["/etc/launchdarkly/overrides.json"])

ldclient.set_config(
    Config(
        "YOUR_SDK_KEY",
        datasystem_config=datasystem.default().overrides(overrides).build(),
    )
)

client = ldclient.get()
```
