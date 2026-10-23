---
id: python-server-sdk/sdk-docs/proxy-env-mac
sdk: python-server-sdk
kind: reference
lang: shell
description: Mac/Linux HTTPS_PROXY environment variable.
# TODO(validate): bare shell export/set commands; not a package-manager install, no harness fits.
# Do not use a run-and-check-exit-0 runner: it would pass bodies that are not correct.
---

```shell
export HTTPS_PROXY=https://web-proxy.domain.com:8080
```
