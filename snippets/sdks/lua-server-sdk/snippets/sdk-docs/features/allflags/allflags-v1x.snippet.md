---
id: lua-server-sdk/sdk-docs/features/allflags/allflags-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: All flags example for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local allFlags = client:allFlags(user)
for flag, value in pairs(allFlags) do
    print(flag, value)
end
--

```
