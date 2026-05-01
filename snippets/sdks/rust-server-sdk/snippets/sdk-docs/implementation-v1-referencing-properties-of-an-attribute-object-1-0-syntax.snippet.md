---
id: rust-server-sdk/sdk-docs/implementation-v1-referencing-properties-of-an-attribute-object-1-0-syntax
sdk: rust-server-sdk
kind: reference
lang: rust
description: "1.0 syntax in section \"Referencing properties of an attribute object\""
---

```rust
let context = ContextBuilder::new("example-context-key")
    .set_value(
        "address",
        AttributeValue::Object(HashMap::from([
            ("street".into(), "Main St".into()),
            ("city".into(), "Springfield".into()),
        ])),
    )
    .build()?;
```
