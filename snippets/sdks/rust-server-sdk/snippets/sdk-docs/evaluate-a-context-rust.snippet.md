---
id: rust-server-sdk/sdk-docs/evaluate-a-context-rust
sdk: rust-server-sdk
kind: reference
lang: rust
description: "Rust in section \"Evaluate a context\""
---

```rust
let context = ContextBuilder::new("example-context-key").build();
let show_feature = client.bool_variation(&context, "example-flag-key", false);

if show_feature {
  # application code to show the feature
} else {
  # the code to run if the feature is off
}
```
