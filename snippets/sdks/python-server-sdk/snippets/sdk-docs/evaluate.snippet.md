---
id: python-server-sdk/sdk-docs/evaluate
sdk: python-server-sdk
kind: reference
lang: python
description: Build a context and evaluate a flag (Python SDK v8.0+).
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
