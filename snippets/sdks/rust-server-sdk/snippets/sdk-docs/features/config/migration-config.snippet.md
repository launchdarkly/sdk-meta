---
id: rust-server-sdk/sdk-docs/features/config/migration-config
sdk: rust-server-sdk
kind: reference
lang: rust
description: Migration configuration example for the Rust SDK v2.2 — read/write methods, execution order, latency/error tracking.

---

```rust
let client = Arc::new(client);
let mut builder = MigratorBuilder::new(client.clone())
    .read(
        |_payload: &String| async move { Ok(()) }.boxed(),
        |_payload: &String| async move { Ok(()) }.boxed(),
        Some(|lhs, rhs| lhs == rhs),
    )
    .write(
        |_payload: &String| async move { Ok(()) }.boxed(),
        |_payload: &String| async move { Ok(()) }.boxed(),
    );

builder = builder
    .read_execution_order(ExecutionOrder::Concurrent) // Or ExecutionOrder::Serial or ExecutionOrder::Random
    .track_latency(true) // defaults to true
    .track_errors(true); // defaults to true

let migrator = builder.build().expect("build migrator");
```
