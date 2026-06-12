---
id: rust-server-sdk/sdk-docs/features/testdata/configure
sdk: rust-server-sdk
kind: reference
lang: rust
description: Test data source configuration for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_server_sdk::{Client, ConfigBuilder, TestData, FlagBuilder};

let td = TestData::new();
// You can set any initial flag states here with td.update

let config = ConfigBuilder::new("YOUR_SDK_KEY")
    .data_source(&td)
    .build()
    .unwrap();
let client = Client::build(config).unwrap();
client.start_with_default_executor();
```
