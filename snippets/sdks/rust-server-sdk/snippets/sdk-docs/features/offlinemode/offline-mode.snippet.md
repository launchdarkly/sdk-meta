---
id: rust-server-sdk/sdk-docs/features/offlinemode/offline-mode
sdk: rust-server-sdk
kind: reference
lang: rust
description: Offline mode example for Rust SDK v1.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
let config = ConfigBuilder::new("YOUR_SDK_KEY").offline(true).build().unwrap();
let ld_client = Client::build(config).unwrap();
ld_client.bool_variation(&context, "example-flag-key", false); // will always return the default value (false)
```
