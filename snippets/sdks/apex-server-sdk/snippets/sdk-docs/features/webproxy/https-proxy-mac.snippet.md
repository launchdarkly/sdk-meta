---
id: apex-server-sdk/sdk-docs/features/webproxy/https-proxy-mac
sdk: apex-server-sdk
kind: reference
lang: bash
description: Mac/Linux HTTPS_PROXY environment variable for the Apex bridge.
# TODO(validate): bare shell export/set commands; not a package-manager install, no harness fits.
# Do not use a run-and-check-exit-0 runner: it would pass bodies that are not correct.
---

```bash
export HTTPS_PROXY=https://my-proxy-host:8080
```
