---
id: python-server-sdk/sdk-docs/features/hooks/define-hook
sdk: python-server-sdk
kind: reference
lang: python
description: Hook implementation and configuration for the Python SDK.
validation:
  scaffold: python-server-sdk/scaffolds/python-syntax-only
---

```python
import ldclient

from ldclient import Config
from ldclient.hook import Hook, Metadata


class ExampleHook(Hook):
    @property
    def metadata(self) -> Metadata:
        return Metadata(name="example-hook")

    # Implement at least one of `before_evaluation`, `after_evaluation`

    # `before_evaluation` is called during the execution of a variation method
    # before the flag value has been determined

    # `after_evaluation` is called during the execution of a variation method
    # after the flag value has been determined


example_hook = ExampleHook()

config = Config("YOUR_SDK_KEY", hooks=[example_hook])

ldclient.set_config(config=config)
client = ldclient.get()
```
