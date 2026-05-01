---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-flagdetailconfig-2-x-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "2.x syntax in section \"Understanding changes to FlagDetailConfig\""
---

```rust
use launchdarkly_server_sdk::FlagDetailConfig;

// Filter for client-side flags only
let config = FlagDetailConfig::new()
    .client_side_only();
```
