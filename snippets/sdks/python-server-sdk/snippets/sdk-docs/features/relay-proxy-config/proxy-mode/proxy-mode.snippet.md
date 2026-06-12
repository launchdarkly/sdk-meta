---
id: python-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: python-server-sdk
kind: reference
lang: python
description: Proxy mode configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
config = Config(sdk_key='YOUR_SDK_KEY',
  stream_uri="https://your-relay-proxy.com:8030",
  base_uri="https://your-relay-proxy.com:8030",
  events_uri="https://your-relay-proxy.com:8030")
```
