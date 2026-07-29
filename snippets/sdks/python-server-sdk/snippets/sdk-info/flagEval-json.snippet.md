---
id: python-server-sdk/sdk-info/flagEval-json
sdk: python-server-sdk
kind: flag-eval
lang: python
file: python-server-sdk/flagEval-json.txt
description: Flag evaluation example for python-server-sdk (JSON flag).
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from ldclient import Context

# Create context using Context builder and use your own values here
context = Context.builder("context-key-123abc").name("Sandy").build()
flag_value = client.variation("featureKey", context, {})

# Use the flag value to configure your feature
# TODO: Put your feature here
```
