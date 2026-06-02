---
id: rust-server-sdk/sdk-docs/features/config/service-endpoint-configuration-eu
sdk: rust-server-sdk
kind: reference
lang: rust
description: Service endpoint configuration example for Rust.
---

```rust
let config = ConfigBuilder::new("YOUR_SDK_KEY").service_endpoints(ServiceEndpointsBuilder::new()
  .streaming_base_url("https://stream.eu.launchdarkly.com")
  .polling_base_url("https://sdk.eu.launchdarkly.com")
  .events_base_url("https://events.eu.launchdarkly.com"));
```
