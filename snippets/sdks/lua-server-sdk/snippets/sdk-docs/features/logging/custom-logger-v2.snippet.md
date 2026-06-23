---
id: lua-server-sdk/sdk-docs/features/logging/custom-logger-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Custom logger configuration example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local logger = ld.makeLogBackend(
    function(level)
        -- Log everything.
        return true
    end,
    function(level, message)
        -- Prints in the format: '[level] hello world'
        print(string.format("[%s] %s", level, message))
    end
)

local config = {
    logging = {
        custom = logger
    }
}
```
