---
id: rust-server-sdk/sdk-docs/features/migrations/stage-switch
sdk: rust-server-sdk
kind: reference
lang: rust
description: Per-stage migration structure for Rust SDK v2.2.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
// define the combination of reads and writes from the new and old systems
// that should occur at each migration stage

match stage {
    Stage::Off => todo!(),
    Stage::DualWrite => todo!(),
    Stage::Live => todo!(),
    Stage::Shadow => todo!(),
    Stage::Rampdown => todo!(),
    Stage::Complete => todo!(),
    _ => todo!(),
};
```
