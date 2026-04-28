---
id: node-client-sdk/getting-started/install
sdk: node-client-sdk
kind: install
lang: shell
description: Install the Node.js client SDK.
inputs:
  version:
    type: string
    description: Optional pinned SDK version; when empty the pin is omitted.
    runtime-default: ""
ld-application:
  slot: install
---

Next, install the LaunchDarkly SDK:

```shell
npm install launchdarkly-node-client-sdk{{ if version }}@{{ version }}{{ end }} --save
```
