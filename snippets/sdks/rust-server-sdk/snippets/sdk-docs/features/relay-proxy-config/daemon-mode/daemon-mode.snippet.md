---
id: rust-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode
sdk: rust-server-sdk
kind: reference
lang: rust
description: Daemon mode configuration example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_server_sdk::{ConfigBuilder, PersistentDataStoreBuilder};


let persistent_store_factory = SomeKindOfFeatureStore::new(store_options);
let persistent_data_store_builder = PersistentDataStoreBuilder::new(Arc::new(persistent_store_factory));

let config = ConfigBuilder::new("sdk-key")
            .daemon_mode(true)
            .data_store(&persistent_data_store_builder)
            .build()
            .expect("config should build");
```
