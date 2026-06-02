---
id: rust-server-sdk/sdk-docs/features/config/service-endpoint-configuration-federal
sdk: rust-server-sdk
kind: reference
lang: rust
description: Service endpoint configuration example for Rust.
---

```rust
let config = ConfigBuilder::new("YOUR_SDK_KEY").service_endpoints(ServiceEndpointsBuilder::new()
  .streaming_base_url("https://stream.launchdarkly.us")
  .polling_base_url("https://sdk.launchdarkly.us")
  .events_base_url("https://events.launchdarkly.us"));
```
