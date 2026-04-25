---
id: python-server-sdk/getting-started/install
sdk: python-server-sdk
kind: install
lang: shell
description: Write the SDK dependency into requirements.txt and install it with pip.
inputs:
  version:
    type: string
    description: Optional pinned SDK version; when empty the pin is omitted.
    runtime-default: ""
ld-application:
  slot: install
---

Create a file called `requirements.txt` with the SDK dependency and install it:

```shell
echo "launchdarkly-server-sdk{{ if version }}=={{ version }}{{ end }}" >> requirements.txt && pip3 install -r requirements.txt
```
