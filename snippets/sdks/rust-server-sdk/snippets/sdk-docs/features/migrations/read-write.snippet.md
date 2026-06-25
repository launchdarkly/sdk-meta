---
id: rust-server-sdk/sdk-docs/features/migrations/read-write
sdk: rust-server-sdk
kind: reference
lang: rust
description: Migration read and write example for Rust SDK v2.2.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
let context = ContextBuilder::new("example-user-key")
    .kind("user")
    .build()
    .expect("Context failed to build");

// this is the migration stage to use if the flag's migration stage
// is not available from LaunchDarkly
let default_stage = Stage::Off;

let read_result = migrator
    .read(
        &context,
        "example-migration-flag-key".into(),
        default_stage,
        "example-payload".into(),
    )
    .await;

let write_result = migrator
    .write(
        &context,
        "example-migration-flag-key".into(),
        default_stage,
        "example-payload".into(),
    )
    .await;
```
