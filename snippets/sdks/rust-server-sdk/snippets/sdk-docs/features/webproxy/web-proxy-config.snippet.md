---
id: rust-server-sdk/sdk-docs/features/webproxy/web-proxy-config
sdk: rust-server-sdk
kind: reference
lang: rust
description: Programmatic web proxy configuration for the Rust SDK transport layer.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
use launchdarkly_sdk_transport::HyperTransport;

// Use a custom proxy URL
let transport = HyperTransport::builder()
    .proxy_url("https://my-proxy-host:8080".to_string())
    .build_http()?;

// Disable proxy completely
let transport = HyperTransport::builder()
    .disable_proxy()
    .build_http()?;
```
