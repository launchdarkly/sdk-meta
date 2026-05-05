---
id: rust-server-sdk/sdk-docs/implementation-v1-understanding-changes-to-private-attributes-1-0-syntax-two-attributes-marked-private-for-all-contexts
sdk: rust-server-sdk
kind: reference
lang: rust
description: "1.0 syntax, two attributes marked private for all contexts in section \"Understanding changes to private attributes\""
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
let config_builder = ConfigBuilder::new("YOUR_SDK_KEY");
let mut processor_builder = EventProcessorBuilder::new();
processor_builder.private_attributes(
    vec!["email".into(), "address".into()]
        .into_iter()
        .collect(),
);
config_builder.event_processor(&processor_builder);
```
