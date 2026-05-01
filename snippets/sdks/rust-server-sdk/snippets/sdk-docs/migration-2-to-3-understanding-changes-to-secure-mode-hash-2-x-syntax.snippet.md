---
id: rust-server-sdk/sdk-docs/migration-2-to-3-understanding-changes-to-secure-mode-hash-2-x-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "2.x syntax in section \"Understanding changes to `secure_mode_hash`\""
---

```rust
use launchdarkly_server_sdk::Client;

let hash: String = client.secure_mode_hash(&context);
```
