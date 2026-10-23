---
id: python-server-sdk/sdk-docs/install
sdk: python-server-sdk
kind: reference
lang: shell
description: pip install for the SDK plus the optional observability plugin (requires Python SDK v9.12+).
# TODO(validate): multi-line pip install body breaks the shell-install harness's per-package extractor.
# To validate, make the harness check one package per line. sdk-info/install-pip covers the SDK install.
---

```shell
pip install launchdarkly-server-sdk

# optional observability plugin, requires Python SDK v9.12+
pip install launchdarkly-observability
```
