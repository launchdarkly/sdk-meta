---
id: rust-server-sdk/sdk-docs/features/datasaving/standard-setup
sdk: rust-server-sdk
kind: reference
lang: rust
description: Data saving mode standard setup for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_server_sdk::{Client, ConfigBuilder, DataSystemBuilder};

let config = ConfigBuilder::new("YOUR_SDK_KEY")
    .data_system(&DataSystemBuilder::default())
    .build()
    .unwrap();

let client = Client::build(config).unwrap();
```
