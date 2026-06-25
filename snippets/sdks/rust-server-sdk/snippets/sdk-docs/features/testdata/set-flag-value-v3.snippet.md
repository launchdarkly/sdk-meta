---
id: rust-server-sdk/sdk-docs/features/testdata/set-flag-value-v3
sdk: rust-server-sdk
kind: reference
lang: rust
description: Setting a test data flag to a specific value for Rust SDK v3.0.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
td.update(FlagBuilder::new("example-flag-key").variation_for_all(false));
```
