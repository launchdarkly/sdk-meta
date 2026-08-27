---
id: rust-server-sdk/sdk-info/init
sdk: rust-server-sdk
kind: init
lang: rust
file: rust-server-sdk/init.txt
description: |
  Client initialization snippet for rust-server-sdk. A complete runnable
  program modeled on launchdarkly/hello-rust (minus env_logger, which the
  batch validator image doesn't pre-bake). The wrappee owns the whole
  program including main, so the harness asserts the snippet's own
  success line via SNIPPET_SUCCESS_RE.
validation:
  runtime: rust
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```rust
use std::time::Duration;

use launchdarkly_server_sdk::{Client, ConfigBuilder};

#[tokio::main]
async fn main() {
    // This is your LaunchDarkly SDK key.
    // Never hardcode your SDK key in production.
    let config = ConfigBuilder::new("YOUR_SDK_KEY")
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
