---
id: lua-server-sdk/sdk-docs/features/storing-data/redis/redis-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Redis source configuration example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
local l = require("launchdarkly_server_sdk")
local r = require("launchdarkly_server_sdk_redis")

local redis = r.makeRedisSource("redis://localhost:6379", "your-key-prefix")
local config = {
    dataSystem = {
        lazyLoad = {
            source = redis
        }
    }
}

local c = l.clientInit("YOUR_SDK_KEY", 1000, config)
```
