---
id: lua-server-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-2-0-syntax-context-with-attributes
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, context with attributes in section \"Working with built-in and custom attributes\""
---

```lua
local context = ld.makeContext({
    user = {
      key       = "example-user-key-1",
      attributes = {
        address   = "123 Main St"
      }
    }
})
```
