---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-cargo-features-webpki-root-certificates
sdk: rust-server-sdk
kind: reference
lang: toml
description: "WebPKI root certificates in section \"Understanding changes to Cargo features\""
---

```toml
[dependencies]
launchdarkly-server-sdk = { version = "3", default-features = false, features = ["hyper-rustls-webpki-roots", "crypto-aws-lc-rs"] }
```
