---
id: lua-server-sdk/sdk-docs/features/relay-proxy-config/proxy-mode/proxy-mode-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Proxy mode configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
local config = {
    serviceEndpoints = {
      streamingBaseURL = "https://your-relay-proxy.com:8030",
      pollingBaseURL = "https://your-relay-proxy.com:8030",
      eventsBaseURL = "https://your-relay-proxy.com:8030"
    }
}
```
