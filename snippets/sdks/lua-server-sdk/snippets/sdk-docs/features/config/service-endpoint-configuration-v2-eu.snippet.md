---
id: lua-server-sdk/sdk-docs/features/config/service-endpoint-configuration-v2-eu
sdk: lua-server-sdk
kind: reference
lang: lua
description: Service endpoint configuration example for Lua.
---

```lua
local config = {
    serviceEndpoints = {
      streamingBaseURL = "https://stream.eu.launchdarkly.com",
      pollingBaseURL = "https://sdk.eu.launchdarkly.com",
      eventsBaseURL = "https://events.eu.launchdarkly.com"
    }
}
```
