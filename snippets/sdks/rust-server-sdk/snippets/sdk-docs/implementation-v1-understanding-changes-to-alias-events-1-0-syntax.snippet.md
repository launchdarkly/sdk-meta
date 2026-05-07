---
id: rust-server-sdk/sdk-docs/implementation-v1-understanding-changes-to-alias-events-1-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "1.0 syntax in section \"Understanding changes to alias events\""
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
let user_context = ContextBuilder::new("example-user-key").build()?;
client.identify(user_context.clone());

let device_context = ContextBuilder::new("example-device-key").kind("device").build()?;
client.identify(device_context.clone());

let multi_context = MultiContextBuilder::new()
    .add_context(user_context)
    .add_context(device_context)
    .build()?;
client.identify(multi_context);
```
