---
id: react-client-sdk/getting-started/legacy-install
sdk: react-client-sdk
kind: install
lang: shell
description: Install the LaunchDarkly React SDK (legacy variant — versioned npm install).
inputs:
  version:
    type: string
    description: Optional pinned SDK version; when empty the pin is omitted.
    runtime-default: ""
ld-application:
  slot: legacy-install
---

Install the LaunchDarkly SDK:

```shell
npm install --save launchdarkly-react-client-sdk{{ if version }}@{{ version }}{{ end }}
```
