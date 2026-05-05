---
id: python-server-sdk/sdk-info/init
sdk: python-server-sdk
kind: init
lang: python
file: python-server-sdk/init.txt
description: Client initialization snippet for python-server-sdk.
validation:
  scaffold: python-server-sdk/scaffolds/init-runner
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
---

```python
import os
import ldclient
from ldclient import Context
from ldclient.config import Config

if __name__ == '__main__':
    # This is your LaunchDarkly SDK key.
    # Never hardcode your SDK key in production.
    ldclient.set_config(Config('YOUR_SDK_KEY'))

    if not ldclient.get().is_initialized():
        print('SDK failed to initialize')
        exit()

    # For onboarding purposes only we flush events as soon as
    # possible so we quickly detect your connection.
    # You don't have to do this in practice because events are automatically flushed.
    ldclient.get().flush()
    print('SDK successfully initialized')
```
