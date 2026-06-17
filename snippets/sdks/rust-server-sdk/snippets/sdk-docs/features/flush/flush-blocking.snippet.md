---
id: rust-server-sdk/sdk-docs/features/flush/flush-blocking
sdk: rust-server-sdk
kind: reference
lang: rust
description: Synchronous event flush example for Rust SDK v3.0.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
use std::time::Duration;

let success = client.flush_blocking(Duration::from_secs(5)).await;
if !success {
    eprintln!("Event flush timed out");
}
```
