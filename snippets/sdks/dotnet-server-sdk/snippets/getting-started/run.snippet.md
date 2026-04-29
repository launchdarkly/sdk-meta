---
id: dotnet-server-sdk/getting-started/run
sdk: dotnet-server-sdk
kind: run
lang: shell
description: Run with dotnet from the command line.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key embedded in the rendered Run command.
ld-application:
  slot: run
---

Run from the command line:

```shell
LAUNCHDARKLY_SDK_KEY={{ apiKey }} dotnet run --project HelloDotNet
```
