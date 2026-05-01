---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-cargo-features-native-tls-backend
sdk: rust-server-sdk
kind: reference
lang: toml
description: "native TLS backend in section \"Understanding changes to Cargo features\""
---

```toml
[dependencies]
launchdarkly-server-sdk = { version = "3", default-features = false, features = ["native-tls", "crypto-aws-lc-rs"] }
```
