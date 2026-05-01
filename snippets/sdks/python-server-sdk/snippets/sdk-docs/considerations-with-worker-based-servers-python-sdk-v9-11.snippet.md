---
id: python-server-sdk/sdk-docs/considerations-with-worker-based-servers-python-sdk-v9-11
sdk: python-server-sdk
kind: reference
lang: python
description: "Python SDK v9.11+ in section \"Considerations with worker-based servers\""
---

```python
# 1. Create the client before forking.
ldclient.set_config(Config("YOUR_SDK_KEY"))
client = ldclient.get()

# 2. From the newly forked process, reinitialize the client by calling `postfork`.
# Examples for specific servers are shown below.
client.postfork()
```
