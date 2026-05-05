---
id: rust-server-sdk/sdk-docs/implementation-v1-understanding-changes-to-private-attributes-1-0-syntax-attribute-marked-private-for-one-context
sdk: rust-server-sdk
kind: reference
lang: rust
description: "1.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
validation:
  scaffold: rust-server-sdk/scaffolds/rust-syntax-only
---

```rust
ContextBuilder::new("example-context-key").add_private_attribute(Reference::new("/address/street"))
```
