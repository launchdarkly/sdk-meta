---
id: rust-server-sdk/sdk-docs/features/config/app-config
sdk: rust-server-sdk
kind: reference
lang: rust
description: Application metadata configuration example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
let mut application_info = ApplicationInfo::new();
application_info
    .application_identifier("authentication-service")
    .application_version("1.0.0");
let config = ConfigBuilder::new(&sdk_key)
    .application_info(application_info)
    .build();
```
