---
id: lua-server-sdk/sdk-docs/features/config/service-endpoint-configuration-v2-relay
sdk: lua-server-sdk
kind: reference
lang: lua
description: Service endpoint configuration example for Lua.
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
