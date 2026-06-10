---
id: rust-server-sdk/sdk-docs/features/evaluation-reasons/evaluation-reasons
sdk: rust-server-sdk
kind: reference
lang: rust
description: Flag evaluation reason example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
let detail = client.bool_variation_detail(&context, "example-flag-key", false);

let value = detail.value;
let index = detail.variation_index;
let reason = detail.reason;
```
