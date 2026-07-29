---
id: python-server-sdk/sdk-info/flagEval-string
sdk: python-server-sdk
kind: flag-eval
lang: python
file: python-server-sdk/flagEval-string.txt
description: Flag evaluation example for python-server-sdk (string flag).
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from ldclient import Context

# Create context using Context builder and use your own values here
context = Context.builder("context-key-123abc").name("Sandy").build()
flag_value = client.variation("featureKey", context, "fallback-value")

if flag_value == "example-variation-value":
    pass  # TODO: Put your feature here
else:
    pass  # TODO: Put your fallback behavior here
```
