---
id: python-server-sdk/ai-configs/initialize
sdk: python-server-sdk
kind: initialize
lang: python
file: python-server-sdk/ai-configs/initialize.txt
description: Initialize the LaunchDarkly client and AI Configs client for python-server-sdk.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
ldclient.set_config(Config("{{sdkkey}}"))
aiclient = LDAIClient(ldclient.get())
```
