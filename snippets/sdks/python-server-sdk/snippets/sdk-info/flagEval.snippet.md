---
id: python-server-sdk/sdk-info/flagEval
sdk: python-server-sdk
kind: flag-eval
lang: python
file: python-server-sdk/flagEval.txt
description: Flag evaluation example for python-server-sdk.
---

```python
from ldclient import Context

# Create context using Context builder and use your own values here
context = Context.builder("context-key-123abc").name("Sandy").build()
flag_value = client.variation("featureKey", context, False)

if flag_value:
    pass  # TODO: Put your feature here
else:
    pass  # TODO: Put your fallback behavior here
```
