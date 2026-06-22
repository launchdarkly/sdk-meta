---
id: rust-server-sdk/sdk-docs/features/privateattrs/config
sdk: rust-server-sdk
kind: reference
lang: rust
description: Private attribute configuration for Rust SDK v1.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only-v1

---

```rust
// All attributes marked private
let config_builder = ConfigBuilder::new("YOUR_SDK_KEY");
let mut processor_builder = EventProcessorBuilder::new();
processor_builder.all_attributes_private(true);
config_builder.event_processor(&processor_builder);

// Two attributes marked private
let config_builder = ConfigBuilder::new("YOUR_SDK_KEY");
let mut processor_builder = EventProcessorBuilder::new();
processor_builder.private_attributes(
    vec![Reference::new("email"), Reference::new("address")]
        .into_iter()
        .collect(),
);
config_builder.event_processor(&processor_builder);
```
