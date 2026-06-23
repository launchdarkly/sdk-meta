---
id: rust-server-sdk/sdk-docs/features/contextconfig/context-kind
sdk: rust-server-sdk
kind: reference
lang: rust
description: Context with a non-user kind for Rust SDK v3.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
let context = ContextBuilder::new("example-context-key")
    .kind("organization")
    .build()?;
```
