---
id: rust-server-sdk/sdk-docs/features/datasaving/relay-proxy-fallback
sdk: rust-server-sdk
kind: reference
lang: rust
description: Data saving mode with Relay Proxy and LaunchDarkly API fallback for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_server_sdk::{
    Client, ConfigBuilder, DataSystemBuilder, FDv2PollingBuilder, FDv2StreamingBuilder,
};
use launchdarkly_sdk_transport::HyperTransport;

let relay_url = "http://my-relay-proxy:8030";

let mut relay_initializer = FDv2PollingBuilder::<HyperTransport>::new();
relay_initializer.base_url(relay_url);

let mut relay_synchronizer = FDv2StreamingBuilder::<HyperTransport>::new();
relay_synchronizer.base_url(relay_url);

let mut data_system = DataSystemBuilder::custom();
data_system
    .initializer(relay_initializer)
    .initializer(FDv2PollingBuilder::<HyperTransport>::new())
    .synchronizer(relay_synchronizer)
    .synchronizer(FDv2StreamingBuilder::<HyperTransport>::new())
    .synchronizer(FDv2PollingBuilder::<HyperTransport>::new());

let config = ConfigBuilder::new("YOUR_SDK_KEY")
    .data_system(&data_system)
    .build()
    .unwrap();

let client = Client::build(config).unwrap();
```
