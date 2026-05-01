---
id: lua-server-sdk/sdk-docs/migration-1-to-2-referencing-properties-of-an-attribute-object-2-0-syntax-context-with-object-attributes
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, context with object attributes in section \"Referencing properties of an attribute object\""
---

```lua
local context = ld.makeContext({
    user = {
      key       = "example-user-key-2",
      attributes = {
        address   = {
          street = "123 Main St",
          city = "Springfield"
        }
      }
    }
})
```
