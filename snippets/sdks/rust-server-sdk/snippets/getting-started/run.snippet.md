---
id: rust-server-sdk/getting-started/run
sdk: rust-server-sdk
kind: run
lang: shell
description: Build and run with cargo.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key embedded in the rendered Run command.
ld-application:
  slot: run
---

Run:

```shell
LAUNCHDARKLY_SDK_KEY='{{ apiKey }}' cargo run
```
