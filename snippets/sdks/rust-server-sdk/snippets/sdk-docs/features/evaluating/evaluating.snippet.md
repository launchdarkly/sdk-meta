---
id: rust-server-sdk/sdk-docs/features/evaluating/evaluating
sdk: rust-server-sdk
kind: reference
lang: rust
description: Flag evaluation example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
let result = client.bool_variation(&context, "your.feature.key", false);
// result is now true or false depending on the setting of this boolean feature flag
```
