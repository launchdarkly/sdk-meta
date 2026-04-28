---
id: rust-server-sdk/getting-started/install-tokio
sdk: rust-server-sdk
kind: install
lang: shell
description: Add tokio with the rt and macros features as a dependency.
ld-application:
  slot: install-tokio
---

Next, add `tokio` as another dependency alongside the SDK.

```shell
cargo add tokio@1 -F rt,macros
```
