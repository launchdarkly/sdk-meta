---
id: lua-server-sdk/sdk-docs/features/privateattrs/context-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Marking context attributes private for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local user = ld.makeContext({
    user = {
        key  = "example-user-key",
        attributes = {
            firstName = "Sandy",
            lastName  = "Smith",
            email     = "sandy@example.com",
            groups    = { "Acme", "Global Health Services" }
        },
        privateAttributes = { "email" }
    }
})
```
