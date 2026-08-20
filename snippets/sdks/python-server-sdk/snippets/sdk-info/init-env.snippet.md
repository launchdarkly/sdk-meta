---
id: python-server-sdk/sdk-info/init-env
sdk: python-server-sdk
kind: init
lang: python
file: python-server-sdk/init-env.txt
description: Client initialization snippet for python-server-sdk reading the SDK key from an environment variable.
validation:
  scaffold: python-server-sdk/scaffolds/init-runner
---

```python
import os
import ldclient
from ldclient import Context
from ldclient.config import Config

if __name__ == '__main__':
    # Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
    # environment variable, so one build can run in every environment.
    ldclient.set_config(Config(os.getenv('LAUNCHDARKLY_SDK_KEY')))

    if not ldclient.get().is_initialized():
        print('SDK failed to initialize')
        exit()

    # For onboarding purposes only we flush events as soon as
    # possible so we quickly detect your connection.
    # You don't have to do this in practice because events are automatically flushed.
    ldclient.get().flush()
    print('SDK successfully initialized')
```
