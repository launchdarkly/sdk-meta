---
id: python-server-sdk/sdk-docs/evaluate-async
sdk: python-server-sdk
kind: reference
lang: python
description: Evaluate a flag with the experimental async client.
validation:
  scaffold: python-server-sdk/scaffolds/python-async-syntax-only
---

```python
from ldclient import Context

context = Context.builder("example-context-key").name("Sandy").build()
flag_value = await client.variation("example-flag-key", context, False)

if flag_value:
    # application code to show the feature
    pass
else:
    # the code to run if the feature is off
    pass
```
