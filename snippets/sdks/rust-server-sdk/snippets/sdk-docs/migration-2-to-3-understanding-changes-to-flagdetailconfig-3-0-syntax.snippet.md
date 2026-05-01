---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-flagdetailconfig-3-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "3.0 syntax in section \"Understanding changes to FlagDetailConfig\""
---

```rust
use launchdarkly_server_sdk::{FlagDetailConfig, FlagFilter};

// Filter for client-side flags only
let config = FlagDetailConfig::new()
    .flag_filter(FlagFilter::CLIENT);

// Filter for mobile flags only
let config = FlagDetailConfig::new()
    .flag_filter(FlagFilter::MOBILE);

// Filter for both client-side and mobile flags
let config = FlagDetailConfig::new()
    .flag_filter(FlagFilter::CLIENT | FlagFilter::MOBILE);
```
