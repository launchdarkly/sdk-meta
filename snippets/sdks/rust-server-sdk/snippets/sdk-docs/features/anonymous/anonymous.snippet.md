---
id: rust-server-sdk/sdk-docs/features/anonymous/anonymous
sdk: rust-server-sdk
kind: reference
lang: rust
description: Anonymous context example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
// Anonymous context with only a key
let context = ContextBuilder::new("example-context-key").anonymous(true).build();

// Anonymous context with a key plus other attributes
let context = ContextBuilder::new("example-context-key").
    anonymous(true).
    set_value("country", "US".into()).
    build();
```
