---
id: rust-server-sdk/sdk-docs/features/contextconfig/multi-context
sdk: rust-server-sdk
kind: reference
lang: rust
description: Multi-context example for Rust SDK v3.
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
