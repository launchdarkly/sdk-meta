---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-secure-mode-hash-3-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "3.0 syntax in section \"Understanding changes to `secure_mode_hash`\""
---

```rust
use launchdarkly_server_sdk::Client;

// Now returns a Result
match client.secure_mode_hash(&context) {
    Ok(hash) => println!("Hash: {}", hash),
    Err(e) => eprintln!("Failed to compute hash: {}", e),
}
```
