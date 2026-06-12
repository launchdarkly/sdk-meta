---
id: rust-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode
sdk: rust-server-sdk
kind: reference
lang: rust
description: Proxy mode configuration example for Rust.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
let config = ConfigBuilder::new("YOUR_SDK_KEY")
    .service_endpoints(
        ServiceEndpointsBuilder::new().relay_proxy("https://your-relay-proxy.com:8030"),
    )
    .build();
```
