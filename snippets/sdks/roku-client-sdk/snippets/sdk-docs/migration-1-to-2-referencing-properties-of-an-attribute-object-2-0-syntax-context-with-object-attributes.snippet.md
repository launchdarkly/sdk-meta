---
id: roku-client-sdk/sdk-docs/migration-1-to-2-referencing-properties-of-an-attribute-object-2-0-syntax-context-with-object-attributes
sdk: roku-client-sdk
kind: reference
lang: brightscript
description: "2.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```brightscript
context = LaunchDarklyCreateContext({
    "kind": "user",
    "key": "example-context-key",
    "address": {"street": "Main St", "city": "Springfield"}
})
```
