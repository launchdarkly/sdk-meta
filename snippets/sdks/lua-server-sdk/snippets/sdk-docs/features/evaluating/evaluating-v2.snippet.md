---
id: lua-server-sdk/sdk-docs/features/evaluating/evaluating-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Flag evaluation example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local value = client:boolVariation(context, "example-flag-key", false)
```
