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
#[allow(dead_code, unused)]
fn _wrappee() {
{{ body }}
}

fn main() {
    println!("feature flag evaluates to true");
}
```
