---
id: lua-server-sdk/sdk-docs/features/config/index-v1
sdk: lua-server-sdk
kind: reference
lang: lua
description: SDK configuration example for Lua.
---

```lua
local config = {
    key                 = "YOUR_SDK_KEY",
    eventsCapacity      = 1000,
    eventsFlushInterval = 30000
}

local client = ld.clientInit(config, 1000)
```
