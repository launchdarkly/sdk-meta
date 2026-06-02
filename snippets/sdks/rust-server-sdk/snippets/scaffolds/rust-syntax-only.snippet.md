---
id: rust-server-sdk/scaffolds/rust-syntax-only
sdk: rust-server-sdk
kind: scaffold
lang: rust
file: src/main.rs
description: |
  Parse-only validator for Rust server SDK doc fragments.
inputs:
  body:
    type: string
    description: The wrappee snippet's rendered body, inserted into the parse-only harness.
validation:
  runtime: rust
  entrypoint: src/main.rs
---

```rust
// Pull SDK types into scope so doc-fragment bodies referencing
// `ContextBuilder`, `MultiContextBuilder`, `Reference`, `Client`, etc.
// resolve at compile time without requiring each fragment to repeat
// the imports.
#[allow(unused_imports)]
use launchdarkly_server_sdk::{
    AttributeValue, Client, ConfigBuilder, Context, ContextBuilder,
    MultiContextBuilder, Reference,
};
#[allow(unused_imports)]
use std::collections::HashMap;

// Stub for the pre-1.0 (beta) `User` API surface — removed at 1.0
// in favor of Context. Doc fragments under
// `implementation-v1-understanding-*-beta-syntax-*` reference the
// old API for side-by-side comparison with the 1.0 equivalents, so
// the syntax-only validator needs a parseable stub to compile them.
#[allow(dead_code, non_camel_case_types)]
struct User;
#[allow(dead_code)]
impl User {
    fn with_key(_key: &str) -> Self { Self }
    fn first_name(self, _v: &str) -> Self { self }
    fn last_name(self, _v: &str) -> Self { self }
    fn email(self, _v: &str) -> Self { self }
    fn custom(self, _k: &str, _v: &str) -> Self { self }
    fn build(self) -> Self { self }
}

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
