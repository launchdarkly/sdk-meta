---
id: lua-server-sdk/sdk-docs/features/relay-proxy-config/daemon-mode/daemon-mode-v1
sdk: lua-server-sdk
kind: reference
lang: lua
description: Daemon mode configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
local backend = makeYourBackendInterface()

local config = {
    key = "YOUR_SDK_KEY",
    featureStoreBackend = backend,
    useLDD = true
}
```
