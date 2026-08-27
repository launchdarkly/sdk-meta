---
id: lua-server-sdk/sdk-info/init
sdk: lua-server-sdk
kind: init
lang: lua
file: lua-server-sdk/init.txt
description: |
  Client initialization snippet for lua-server-sdk. A complete runnable
  program modeled on lua-server-sdk's examples/hello-lua-server. The
  wrappee owns the whole program, so the harness asserts the snippet's
  own success line via SNIPPET_SUCCESS_RE. clientInit blocks while the
  SDK connects, so reaching the print proves the FFI init completed.
validation:
  runtime: lua-server
  placeholders:
    YOUR_SDK_KEY: LAUNCHDARKLY_SDK_KEY
  env:
    SNIPPET_SUCCESS_RE: SDK successfully initialized
---

```lua
local ld = require("launchdarkly_server_sdk")

local config = {}

-- This is your LaunchDarkly SDK key.
-- Never hardcode your SDK key in production.
-- clientInit blocks for up to 1 second while the SDK connects.
local client = ld.clientInit("YOUR_SDK_KEY", 1000, config)

print("SDK successfully initialized")
```
