---
id: rust-server-sdk/sdk-docs/features/config/index
sdk: rust-server-sdk
kind: reference
lang: rust
description: SDK configuration example for Rust.
---

```rust
let config = ConfigBuilder::new("YOUR_SDK_KEY")
  .offline(true)
  .build();
let client = Client::build(config).unwrap();
```
