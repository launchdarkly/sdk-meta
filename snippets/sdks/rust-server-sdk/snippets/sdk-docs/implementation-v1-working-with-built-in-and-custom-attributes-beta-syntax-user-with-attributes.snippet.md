---
id: rust-server-sdk/sdk-docs/implementation-v1-working-with-built-in-and-custom-attributes-beta-syntax-user-with-attributes
sdk: rust-server-sdk
kind: reference
lang: rust
description: "Beta syntax, user with attributes in section \"Working with built-in and custom attributes\""
# TODO(validator): body's `.into()` on `"groups"` and on
# `vec!["Acme", ...]` requires type inference toward
# `AttributeValue` (the pre-1.0 beta SDK's variant type). The
# rust-syntax-only scaffold's `User::custom<T>` generic doesn't
# narrow T enough for the inference to converge — bytes::Bytes and
# hyper::body::Body each implement From<&'static str>, so rustc
# picks neither. Fix by either binding to a pre-1.0 rust SDK
# version-pinned scaffold or by narrowing `User::custom` to take a
# concrete `HashMap<String, AttributeValue>` once the AttributeValue
# stub is also added.
---

```rust
// User with only a key
let user1 = User::with_key("example-user-key").build();

// User with a key plus other attributes
let custom = hashmap! {
    "groups".into() => vec!["Acme", "Global Health Services"].into(),
};
let mut builder = User::with_key("user-key-456def");
builder.first_name("Sandy");
builder.last_name("Smith");
builder.email("sandy@example.com");
builder.custom(custom);
let user2 = builder.build();
```
