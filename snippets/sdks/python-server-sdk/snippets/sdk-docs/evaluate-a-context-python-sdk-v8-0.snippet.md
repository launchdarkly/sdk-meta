---
id: python-server-sdk/sdk-docs/evaluate-a-context-python-sdk-v8-0
sdk: python-server-sdk
kind: reference
lang: python
description: "Python SDK v8.0+ in section \"Evaluate a context\""
---

```python
from ldclient import Context

context = Context.builder("example-context-key").name("Sandy").build()
flag_value = client.variation("example-flag-key", context, False)

if flag_value:
    # application code to show the feature
else:
    # the code to run if the feature is off
```
