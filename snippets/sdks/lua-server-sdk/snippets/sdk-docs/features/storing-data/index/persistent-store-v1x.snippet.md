---
id: lua-server-sdk/sdk-docs/features/storing-data/index/persistent-store-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: Persistent feature store configuration example for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local l = require("launchdarkly_server_sdk")

local backend = makeYourBackendInterface()

local c = l.clientInit({
    key                 = "YOUR_SDK_KEY",
    featureStoreBackend = backend
})
```
