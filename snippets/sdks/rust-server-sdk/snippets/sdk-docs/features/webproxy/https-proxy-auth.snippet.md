---
id: rust-server-sdk/sdk-docs/features/webproxy/https-proxy-auth
sdk: rust-server-sdk
kind: reference
lang: bash
description: HTTPS_PROXY environment variable with credentials for the Rust SDK.
# TODO(validate): bare shell export/set commands; not a package-manager install, no harness fits.
# Do not use a run-and-check-exit-0 runner: it would pass bodies that are not correct.
---

```bash
export HTTPS_PROXY="https://username:password@my-proxy-host:8080"
```
