---
id: python-server-sdk/sdk-docs/features/webproxy/https-proxy-windows
sdk: python-server-sdk
kind: reference
lang: bash
description: Windows HTTPS_PROXY environment variable for the Python SDK.
# TODO(validate): bare shell export/set commands; not a package-manager install, no harness fits.
# Do not use a run-and-check-exit-0 runner: it would pass bodies that are not correct.
---

```bash
set HTTPS_PROXY=https://my-proxy-host:8080
```
