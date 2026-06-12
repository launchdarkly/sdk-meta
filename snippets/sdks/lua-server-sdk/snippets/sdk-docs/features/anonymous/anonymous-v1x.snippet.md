---
id: lua-server-sdk/sdk-docs/features/anonymous/anonymous-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: Anonymous user example for Lua, SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
local user = ld.makeUser({
    key       = "example-user-key",
    anonymous = true
    }
)
```
