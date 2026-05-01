---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-the-new-transport-layer-2-x-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "2.x syntax in section \"Understanding the new transport layer\""
---

```rust
use launchdarkly_server_sdk::ConfigBuilder;
use hyper_rustls::HttpsConnectorBuilder;

let connector = HttpsConnectorBuilder::new()
    .with_native_roots()
    .https_or_http()
    .enable_http1()
    .build();

let config = ConfigBuilder::new("sdk-key")
    .datasource_with_connector(&connector)
    .build();
```
