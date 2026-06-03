---
id: rust-server-sdk/sdk-docs/implementation-v1-understanding-changes-to-private-attributes-1-0-syntax-two-attributes-marked-private-for-all-contexts
sdk: rust-server-sdk
kind: reference
lang: rust
description: "1.0 syntax, two attributes marked private for all contexts in section \"Understanding changes to private attributes\""
# TODO(scaffold): body uses EventProcessorBuilder::new() which needs
# explicit type-parameter annotation in the current rust SDK
# (E0282/E0283 — type annotations needed for `EventProcessorBuilder<_>`).
# Either: add a type alias to the rust scaffold pinning the default
# factory type, or pin a v1.0-era rust SDK in a parallel
# rust-syntax-only-v1 scaffold (this snippet's name explicitly
# documents the 1.0 syntax).
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
