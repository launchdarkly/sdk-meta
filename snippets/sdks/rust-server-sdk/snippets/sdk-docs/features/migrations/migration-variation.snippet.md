---
id: rust-server-sdk/sdk-docs/features/migrations/migration-variation
sdk: rust-server-sdk
kind: reference
lang: rust
description: Migration stage evaluation (migration_variation) for Rust SDK v2.2.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
let context = ContextBuilder::new("example-user-key")
    .kind("user")
    .build()
    .expect("Context failed to build");

let (stage, tracker) =
    client.migration_variation(&context, "example-migration-flag-key", Stage::Off);
```
