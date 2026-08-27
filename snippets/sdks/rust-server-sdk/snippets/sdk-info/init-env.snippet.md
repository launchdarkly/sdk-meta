---
id: rust-server-sdk/sdk-info/init-env
sdk: rust-server-sdk
kind: init
lang: rust
file: rust-server-sdk/init-env.txt
description: Client initialization snippet for rust-server-sdk reading the SDK key from an environment variable.
validation:
  runtime: rust
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```rust
use std::time::Duration;

use launchdarkly_server_sdk::{Client, ConfigBuilder};

#[tokio::main]
async fn main() {
    // Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
    // environment variable, so one build can run in every environment.
    let sdk_key = std::env::var("LAUNCHDARKLY_SDK_KEY")
        .expect("LAUNCHDARKLY_SDK_KEY env should be set");
    let config = ConfigBuilder::new(&sdk_key)
        .build()
        .expect("config failed to build");
    let client = Client::build(config).expect("client failed to build");

    // Starts the client using the currently active runtime.
    client.start_with_default_executor();

    // Wait up to five seconds for the client to connect to LaunchDarkly.
    let initialized = client
        .wait_for_initialization(Duration::from_secs(5))
        .await
        .unwrap_or(false);
    if !initialized {
        println!("SDK failed to initialize");
        std::process::exit(1);
    }

    println!("SDK successfully initialized");
}
```
