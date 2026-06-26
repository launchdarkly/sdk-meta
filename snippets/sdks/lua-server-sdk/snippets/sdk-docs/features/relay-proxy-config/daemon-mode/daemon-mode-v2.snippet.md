---
id: lua-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Daemon mode (lazy load) configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
local config = {
    dataSystem = {
        lazyLoad = {
            source = makeYourSource()
        }
    }
}
```
