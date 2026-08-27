---
id: haskell-server-sdk/sdk-info/install-stack
sdk: haskell-server-sdk
kind: install
lang: yaml
file: haskell-server-sdk/install-stack.txt
description: |
  Install fragment for haskell-server-sdk (stack). The SDK is published to
  Hackage but not to Stackage snapshots, so stack projects need the
  package listed in `extra-deps` as well as in the package dependencies.
  Mirrors launchdarkly/hello-haskell-server's stack.yaml + package.yaml.
---

```yaml
# stack.yaml
extra-deps:
- launchdarkly-server-sdk-4.6.0

# package.yaml
dependencies:
- launchdarkly-server-sdk
```
