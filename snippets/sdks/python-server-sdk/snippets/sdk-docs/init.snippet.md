---
id: python-server-sdk/sdk-docs/init
sdk: python-server-sdk
kind: reference
lang: python
description: Initialize the singleton ldclient with the SDK key and the optional observability plugin.
---

```python
ldclient.set_config(Config("YOUR_SDK_KEY",
  # optional observability plugin, requires Python SDK v9.12+
  plugins=[ObservabilityPlugin()]
  )
)
client = ldclient.get()
```
