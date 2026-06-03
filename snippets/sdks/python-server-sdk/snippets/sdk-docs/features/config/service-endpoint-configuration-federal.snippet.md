---
id: python-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: python-server-sdk
kind: reference
lang: python
description: Service endpoint configuration example for Python.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only

---

```python
config = Config(sdk_key='YOUR_SDK_KEY',
  stream_uri="https://stream.launchdarkly.us",
  base_uri="https://sdk.launchdarkly.us",
  events_uri="https://events.launchdarkly.us")
```
