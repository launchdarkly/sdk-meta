---
id: node-server-sdk/getting-started/install
sdk: node-server-sdk
kind: install
lang: shell
description: Install the Node.js server SDK.
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
npm install @launchdarkly/node-server-sdk{{ if version }}@{{ version }}{{ end }} --save
```
