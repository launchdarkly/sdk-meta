---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-event-processor-configuration-3-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "3.0 syntax in section \"Understanding changes to event processor configuration\""
---

```rust
use launchdarkly_server_sdk::{ConfigBuilder, EventProcessorBuilder};
use launchdarkly_sdk_transport::HyperTransport;

let transport = HyperTransport::new(my_connector);

let builder = EventProcessorBuilder::new()
    .transport(transport);
    // compress_events is now true by default

let config = ConfigBuilder::new("sdk-key")
    .event_processor(builder)
    .build();
```
