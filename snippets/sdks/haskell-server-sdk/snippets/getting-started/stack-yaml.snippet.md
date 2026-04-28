---
id: haskell-server-sdk/getting-started/stack-yaml
sdk: haskell-server-sdk
kind: manifest-fragment
lang: yaml
description: extra-deps entry to add to stack.yaml.
inputs:
  version:
    type: string
    description: SDK version. Gonfalon fetches the latest from Hackage asynchronously.
    runtime-default: ""
ld-application:
  slot: stack-yaml
---

Add the SDK version as an `extra-deps` entry in `stack.yaml`:

```yaml
- launchdarkly-server-sdk-{{ version }}
```
