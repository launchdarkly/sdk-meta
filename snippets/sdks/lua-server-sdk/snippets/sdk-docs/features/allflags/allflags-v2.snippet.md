---
id: lua-server-sdk/sdk-docs/features/allflags/allflags-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: All flags example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local allFlags = client:allFlags(context)
for flag, value in pairs(allFlags) do
    print(flag, value)
end
```
