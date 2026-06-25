---
id: rust-server-sdk/sdk-docs/features/testdata/flag-behavior-v3
sdk: rust-server-sdk
kind: reference
lang: rust
description: Configuring test data flag behavior for Rust SDK v3.0.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_server_sdk::{AttributeValue, FlagValue, Kind};
use std::convert::TryFrom;

// This flag is true for the context with the key "example-context-key" and kind of "organization",
// and false for everyone else
td.update(FlagBuilder::new("flag-key-456def")
    .variation_for_key(Kind::try_from("organization").unwrap(), "example-context-key", true)
    .fallthrough_variation(false));

// This flag returns the string variation "green" for contexts that have the
// attribute "admin" with a value of true, and "red" for everyone else.
td.update(FlagBuilder::new("flag-key-789ghi")
    .variations(vec![FlagValue::Str("red".into()), FlagValue::Str("green".into())])
    .fallthrough_variation_index(0)
    .if_match("admin", vec![AttributeValue::Bool(true)])
    .then_return_index(1));
```
