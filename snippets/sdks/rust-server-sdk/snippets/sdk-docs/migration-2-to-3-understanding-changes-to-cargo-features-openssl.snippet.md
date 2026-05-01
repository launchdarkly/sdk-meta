---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-cargo-features-openssl
sdk: rust-server-sdk
kind: reference
lang: toml
description: "OpenSSL in section \"Understanding changes to Cargo features\""
---

```toml
[dependencies]
launchdarkly-server-sdk = { version = "3", default-features = false, features = ["hyper-rustls-native-roots", "crypto-openssl"] }
```
