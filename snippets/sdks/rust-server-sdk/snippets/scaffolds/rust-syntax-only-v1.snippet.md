---
id: rust-server-sdk/scaffolds/rust-syntax-only-v1
sdk: rust-server-sdk
kind: scaffold
lang: rust
file: src/main.rs
description: |
  Parse-only validator for Rust server SDK doc fragments that target
  the v1.x API surface. The v1-era `EventProcessorBuilder` is
  non-generic, so fragments that say `EventProcessorBuilder::new()`
  compile against 1.x but not against 2.x (where the builder gained a
  connector type parameter with no defaulted inference in expression
  position).

  Routes through the same `rust` validator container as
  `rust-syntax-only`, but sets `LD_RUST_SDK_VERSION=1` via
  `validation.env` so the harness's `cargo add` resolves the latest
  1.x release instead of the latest overall.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: rust
  entrypoint: src/main.rs
  env:
    LD_RUST_SDK_VERSION: "1"
---

```rust
// Pull SDK types into scope so doc-fragment bodies referencing
// `ContextBuilder`, `Reference`, `EventProcessorBuilder`, etc.
// resolve at compile time without requiring each fragment to repeat
// the imports. Restricted to names the 1.x release exports.
#[allow(unused_imports)]
use launchdarkly_server_sdk::{
    ApplicationInfo, AttributeValue, Client, ConfigBuilder, Context, ContextBuilder,
    EventProcessorBuilder, MultiContextBuilder, Reason, Reference, ServiceEndpointsBuilder,
};
#[allow(unused_imports)]
use std::sync::Arc;
#[allow(unused_imports)]
use std::collections::HashMap;

// Placeholder constants the docs reference directly (`YOUR_SDK_KEY`)
// rather than substituting at render time.
#[allow(non_upper_case_globals, dead_code)]
const YOUR_SDK_KEY: &str = "";

#[allow(dead_code, unused, unused_variables, unused_must_use)]
async fn _wrappee() -> Result<(), Box<dyn std::error::Error>> {
    let client: Client = unimplemented!();
    let context = ContextBuilder::new("stub").build()?;
{{ body }}
    Ok(())
}

#[tokio::main]
async fn main() {
    println!("feature flag evaluates to true");
}
```
