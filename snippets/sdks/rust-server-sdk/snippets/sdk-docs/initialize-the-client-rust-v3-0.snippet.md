---
id: rust-server-sdk/sdk-docs/initialize-the-client-rust-v3-0
sdk: rust-server-sdk
kind: reference
lang: rust
description: "Rust, v3.0+ in section \"Initialize the client\""
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use std::time::Duration;

#[tokio::main]
async fn main() {
    let config = ConfigBuilder::new(&YOUR_SDK_KEY).build().unwrap();
    let client = Client::build(config).unwrap();

    client.start_with_default_executor();

    let initialized = client
        .wait_for_initialization(Duration::from_secs(10))
        .await
        .unwrap_or(false);

    if !initialized {
        panic!("Client failed to successfully initialize");
    }
}
```
