---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-data-source-configuration-3-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "3.0 syntax in section \"Understanding changes to data source configuration\""
---

```rust
use launchdarkly_server_sdk::{ConfigBuilder, StreamingDataSourceBuilder};
use launchdarkly_sdk_transport::HyperTransport;

let transport = HyperTransport::new(my_connector);

let builder = StreamingDataSourceBuilder::new()
    .transport(transport);

let config = ConfigBuilder::new("sdk-key")
    .datasource(builder)
    .build();
```
