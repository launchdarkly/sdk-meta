---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-the-new-transport-layer-3-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "3.0 syntax in section \"Understanding the new transport layer\""
---

```rust
use launchdarkly_server_sdk::ConfigBuilder;
use launchdarkly_sdk_transport::HyperTransport;
use hyper_rustls::HttpsConnectorBuilder;

let connector = HttpsConnectorBuilder::new()
    .with_native_roots()
    .unwrap()
    .https_or_http()
    .enable_http1()
    .build();

let transport = HyperTransport::new(connector);

let config = ConfigBuilder::new("sdk-key")
    .datasource_with_transport(&transport)
    .build();
```
