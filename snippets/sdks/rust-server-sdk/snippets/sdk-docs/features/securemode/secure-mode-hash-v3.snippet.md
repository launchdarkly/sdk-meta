---
id: rust-server-sdk/sdk-docs/features/securemode/secure-mode-hash-v3
sdk: rust-server-sdk
kind: reference
lang: rust
description: Secure mode hash example for Rust SDK v3.0.
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only

---

```rust
match client.secure_mode_hash(&context) {
    Ok(hash) => println!("Hash: {}", hash),
    Err(e) => eprintln!("Failed to compute hash: {}", e),
}
```
