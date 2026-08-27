---
id: lua-server-sdk/sdk-info/init-env
sdk: lua-server-sdk
kind: init
lang: lua
file: lua-server-sdk/init-env.txt
description: Client initialization snippet for lua-server-sdk reading the SDK key from an environment variable.
validation:
  runtime: lua-server
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```lua
local ld = require("launchdarkly_server_sdk")

local config = {}

-- Reads your LaunchDarkly SDK key from the LAUNCHDARKLY_SDK_KEY
-- environment variable, so one build can run in every environment.
-- clientInit blocks for up to 1 second while the SDK connects.
local client = ld.clientInit(os.getenv("LAUNCHDARKLY_SDK_KEY"), 1000, config)

print("SDK successfully initialized")
```
