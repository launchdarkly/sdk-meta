---
id: lua-server-sdk/sdk-docs/features/contextconfig/context-example-v1
sdk: lua-server-sdk
kind: reference
lang: lua
description: User example for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local user = ld.makeUser({
    key       = "example-user-key",
    firstName = "Sandy",
    lastName  = "Smith",
    email     = "sandy@example.com",
    custom    = {
        groups = { "Acme", "Global Health Services" }
    }
})
```
