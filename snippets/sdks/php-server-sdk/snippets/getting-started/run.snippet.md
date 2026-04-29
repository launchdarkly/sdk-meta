---
id: php-server-sdk/getting-started/run
sdk: php-server-sdk
kind: run
lang: shell
description: Run the program with the SDK key in the environment.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key embedded in the rendered Run command.
ld-application:
  slot: run
---

Run:

```shell
LAUNCHDARKLY_SDK_KEY='{{ apiKey }}' php main.php
```
