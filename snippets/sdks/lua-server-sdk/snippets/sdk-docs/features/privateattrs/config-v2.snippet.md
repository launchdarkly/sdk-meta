---
id: lua-server-sdk/sdk-docs/features/privateattrs/config-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Private attribute configuration for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
-- sets all attributes private
local configAllPrivate = {
    events = {
        allAttributesPrivate = true
    }
}

-- sets "email" and "address" private
local configSomePrivate = {
    events = {
        privateAttributes = { "email", "address" }
    }
}
```
