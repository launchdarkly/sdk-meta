---
id: lua-server-sdk/sdk-docs/features/storing-data/redis/redis-v1x
sdk: lua-server-sdk
kind: reference
lang: lua
description: Redis feature store configuration example for Lua SDK v1.x.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local l = require("launchdarkly_server_sdk")
local r = require("launchdarkly_server_sdk_redis")

local backend = r.makeStore({
    host   = "your-redis",
    port   = 6379,
    prefix = "your-key-prefix"
})

local c = l.clientInit({
    key                 = "YOUR_SDK_KEY",
    featureStoreBackend = backend
})
```
