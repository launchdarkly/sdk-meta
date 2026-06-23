---
id: rust-server-sdk/sdk-docs/features/privateattrs/context
sdk: rust-server-sdk
kind: reference
lang: rust
description: Marking context attributes private with the context builder for Rust SDK v1.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only-v1

---

```rust
let context = ContextBuilder::new("example-context-key")
    .set_value("email", "youremail@example.com".into())
    .add_private_attribute(Reference::new("email"))
    .build()?;
```
