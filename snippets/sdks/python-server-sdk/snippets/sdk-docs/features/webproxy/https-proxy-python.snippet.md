---
id: python-server-sdk/sdk-docs/features/webproxy/https-proxy-python
sdk: python-server-sdk
kind: reference
lang: python
description: Set HTTPS_PROXY from inside Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import os

os.environ["HTTPS_PROXY"] = "https://my-proxy-host:8080"
```
