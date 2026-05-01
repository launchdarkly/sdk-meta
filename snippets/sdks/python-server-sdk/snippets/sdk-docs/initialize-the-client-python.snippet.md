---
id: python-server-sdk/sdk-docs/initialize-the-client-python
sdk: python-server-sdk
kind: reference
lang: python
description: "Python in section \"Initialize the client\""
---

```python
ldclient.set_config(Config("YOUR_SDK_KEY",
  # optional observability plugin, requires Python SDK v9.12+
  plugins=[ObservabilityPlugin()]
  )
)
client = ldclient.get()
```
