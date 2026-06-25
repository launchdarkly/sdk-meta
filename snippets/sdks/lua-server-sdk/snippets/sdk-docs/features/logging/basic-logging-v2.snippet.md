---
id: lua-server-sdk/sdk-docs/features/logging/basic-logging-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Default-logger configuration example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local config = {
    logging = {
        basic = {
            tag = "launchdarkly",
            level = "warn"
        }
    }
}
```
