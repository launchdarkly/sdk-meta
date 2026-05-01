---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-data-source-configuration-2-x-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "2.x syntax in section \"Understanding changes to data source configuration\""
---

```rust
use launchdarkly_server_sdk::{ConfigBuilder, StreamingDataSourceBuilder};

let builder = StreamingDataSourceBuilder::new()
    .https_connector(my_connector);

let config = ConfigBuilder::new("sdk-key")
    .datasource(builder)
    .build();
```
