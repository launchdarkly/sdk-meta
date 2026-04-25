---
id: python-server-sdk/getting-started/run
sdk: python-server-sdk
kind: run
lang: shell
description: Run the hello-world program with the LaunchDarkly SDK key in the environment.
inputs:
  apiKey:
    type: sdk-key
    description: The SDK key to embed in the rendered Run command.
ld-application:
  slot: run
---

Run the program:

```shell
LAUNCHDARKLY_SDK_KEY="{{ apiKey }}" python main.py
```
