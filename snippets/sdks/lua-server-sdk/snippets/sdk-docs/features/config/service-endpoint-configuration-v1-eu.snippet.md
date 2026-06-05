---
id: lua-server-sdk/sdk-docs/features/config/service-endpoint-configuration-v1-eu
sdk: lua-server-sdk
kind: reference
lang: lua
description: Service endpoint configuration example for Lua.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local config = {
    key = "YOUR_SDK_KEY",
    streamURI = "https://stream.eu.launchdarkly.com",
    baseURI = "https://sdk.eu.launchdarkly.com",
    eventsURI = "https://events.eu.launchdarkly.com",
}
```
