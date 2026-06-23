---
id: rust-server-sdk/scaffolds/rust-syntax-only-v2
sdk: rust-server-sdk
kind: scaffold
lang: rust
file: src/main.rs
description: |
  Parse-only validator for Rust server SDK v2.x-era doc fragments.

  The rust validator's Cargo project pulls the latest crate (3.x),
  where some client methods changed shape — `secure_mode_hash` returned
  `String` in 2.x but returns `Result<String, String>` in 3.0. v2-era
  fragments therefore compile against a stub client that mirrors the
  2.x surface instead of the real `Client`, following the same approach
  as the beta `User` stub in `rust-syntax-only`. Extend the stub's
  method set as future v2-era fragments need it.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: rust
  entrypoint: src/main.rs
---

```rust
#[allow(unused_imports)]
use launchdarkly_server_sdk::{Context, ContextBuilder};

// Stub of the v2-era `Client` API surface. Never executed; present so
// v2-era bodies type-check without the real (3.x) client in scope.
#[allow(dead_code)]
struct V2Client;
#[allow(dead_code)]
impl V2Client {
    fn secure_mode_hash(&self, _context: &Context) -> String {
        String::new()
    }
}

#[allow(dead_code, unused, unused_variables, unused_must_use)]
async fn _wrappee() -> Result<(), Box<dyn std::error::Error>> {
    let client = V2Client;
    let context = ContextBuilder::new("stub").build()?;
{{ body }}
    Ok(())
}

#[tokio::main]
async fn main() {
    println!("feature flag evaluates to true");
}
```
