---
id: haskell-server-sdk/getting-started/run
sdk: haskell-server-sdk
kind: run
lang: shell
description: Build with stack and run.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key embedded in the rendered Run command.
ld-application:
  slot: run
---

Build and run:

```shell
stack build && LAUNCHDARKLY_SDK_KEY='{{ apiKey }}' stack exec hello-haskell-exe
```
