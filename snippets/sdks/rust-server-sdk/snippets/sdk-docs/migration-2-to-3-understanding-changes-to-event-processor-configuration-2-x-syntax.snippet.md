---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-event-processor-configuration-2-x-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "2.x syntax in section \"Understanding changes to event processor configuration\""
---

```rust
use launchdarkly_server_sdk::{ConfigBuilder, EventProcessorBuilder};

let builder = EventProcessorBuilder::new()
    .https_connector(my_connector)
    .compress_events(true); // had to opt in

let config = ConfigBuilder::new("sdk-key")
    .event_processor(builder)
    .build();
```
