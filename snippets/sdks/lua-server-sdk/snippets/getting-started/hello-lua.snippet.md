---
id: lua-server-sdk/getting-started/hello-lua
sdk: lua-server-sdk
kind: hello-world
lang: lua
file: hello.lua
description: Hello-world program that initializes the Lua server SDK and evaluates a feature flag.
inputs:
  apiKey:
    type: sdk-key
    description: SDK key baked into the rendered source.
  featureKey:
    type: flag-key
    description: Default flag key baked into the rendered source.
ld-application:
  slot: hello-lua
validation:
  runtime: lua-server
  entrypoint: hello.lua
---

Create a file named `hello.lua` and add the following code:

```lua
local ld = require("launchdarkly_server_sdk")
local config = {}

local client = ld.clientInit("{{ apiKey }}", 1000, config)

local user = ld.makeContext({
    user = {
        key = "example-user-key",
        name = "Sandy"
    }
})

local value = client:boolVariation(user, "{{ featureKey }}", false)
print("*** The '{{ featureKey }}' feature flag evaluates to "..tostring(value)..".")
```
