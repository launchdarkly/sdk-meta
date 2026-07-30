---
id: python-server-sdk/sdk-info/flagEval-number
sdk: python-server-sdk
kind: flag-eval
lang: python
file: python-server-sdk/flagEval-number.txt
description: Flag evaluation example for python-server-sdk (number flag).
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
from ldclient import Context

# Create context using Context builder and use your own values here
context = Context.builder("context-key-123abc").name("Sandy").build()
flag_value = client.variation("featureKey", context, 0)

if flag_value == 1:
    pass  # TODO: Put your feature here
else:
    pass  # TODO: Put your fallback behavior here
```
