---
id: lua-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-v1
sdk: lua-server-sdk
kind: reference
lang: lua
description: Proxy mode configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
local config = {
    key = "YOUR_SDK_KEY",
    streamURI = "https://your-relay-proxy.com:8030",
    baseURI = "https://your-relay-proxy.com:8030",
    eventsURI = "https://your-relay-proxy.com:8030",
}
```
