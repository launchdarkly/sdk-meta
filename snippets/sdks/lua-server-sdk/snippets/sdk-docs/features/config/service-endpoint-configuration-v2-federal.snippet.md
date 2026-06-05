---
id: lua-server-sdk/sdk-docs/features/config/service-endpoint-configuration-v2-federal
sdk: lua-server-sdk
kind: reference
lang: lua
description: Service endpoint configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local config = {
    serviceEndpoints = {
      streamingBaseURL = "https://stream.launchdarkly.us",
      pollingBaseURL = "https://sdk.launchdarkly.us",
      eventsBaseURL = "https://events.launchdarkly.us"
    }
}
```
