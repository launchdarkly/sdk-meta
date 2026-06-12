---
id: lua-server-sdk/sdk-docs/features/privateattrs/config-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: Private attribute configuration for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
-- sets all attributes private
local configAllPrivate = {
    allAttributesPrivate = true
}

-- sets "email" and "address" private
local configSomePrivate = {
    privateAttributeNames = { "email", "address" }
}
```
