---
id: lua-server-sdk/sdk-docs/features/config/index-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: SDK configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local config = {
    events = {
        capacity = 1000,
        flushIntervalMilliseconds  = 30000
    }
}

-- This blocks for 1 second to initialize
local client = ld.clientInit("YOUR_SDK_KEY", 1000, config)
```
