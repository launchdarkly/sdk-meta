---
id: rust-server-sdk/sdk-docs/features/contextconfig/context-example
sdk: rust-server-sdk
kind: reference
lang: rust
description: Context example for Rust SDK v3.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
// Context with only a key
let context = ContextBuilder::new("example-context-key").build()?;

// Context with a key plus other attributes
let context = ContextBuilder::new("example-context-key")
    .set_value("first_name", "Sandy".into())
    .set_value("last_name", "Smith".into())
    .set_value("email", "sandy@example.com".into())
    .set_value("groups", vec!["Acme", "Global Health Services"].into())
    .build();
```
