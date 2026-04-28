---
id: go-server-sdk/getting-started/run
sdk: go-server-sdk
kind: run
lang: shell
description: Build and run the program with the SDK key in the environment.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key to embed in the rendered Run command.
ld-application:
  slot: run
---

Build and run:

```shell
go build && LAUNCHDARKLY_SDK_KEY='{{ apiKey }}' ./hello-go
```
