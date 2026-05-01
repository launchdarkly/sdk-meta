---
id: lua-server-sdk/sdk-docs/migration-1-to-2-working-with-built-in-and-custom-attributes-1-x-syntax-user-with-attributes
sdk: lua-server-sdk
kind: reference
lang: lua
description: "1.x syntax, user with attributes in section \"Working with built-in and custom attributes\""
---

```lua
local user = ld.makeUser({
    key       = "example-user-key",
    custom    = {
        address = { "123 Main St" }
    }
})
```
