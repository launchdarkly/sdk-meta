---
id: haskell-server-sdk/getting-started/package-yaml
sdk: haskell-server-sdk
kind: manifest-fragment
lang: yaml
description: Dependencies entry to add to package.yaml.
ld-application:
  slot: package-yaml
---

Next, add the SDK and `text` package to your list of dependencies in `package.yaml`:

```yaml
launchdarkly-server-sdk, text
```
