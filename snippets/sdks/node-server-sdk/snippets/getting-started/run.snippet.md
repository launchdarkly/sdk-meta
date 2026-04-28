---
id: node-server-sdk/getting-started/run
sdk: node-server-sdk
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
LAUNCHDARKLY_SDK_KEY={{ apiKey }} node index.js
```
