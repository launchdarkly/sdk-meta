---
id: python-server-sdk/sdk-docs/features/config/service-endpoint-configuration-relay
sdk: python-server-sdk
kind: reference
lang: python
description: Service endpoint configuration example for Python.
---

```python
config = Config(sdk_key='YOUR_SDK_KEY',
  stream_uri="https://your-relay-proxy.com:8030",
  base_uri="https://your-relay-proxy.com:8030",
  events_uri="https://your-relay-proxy.com:8030")
```
