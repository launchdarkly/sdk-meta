---
id: lua-server-sdk/sdk-docs/features/anonymous/anonymous-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Anonymous context example for Lua, SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only
---

```lua
-- to create an anonymous user context
local userContext = ld.makeContext({
    user = {
      key       = "example-user-key",
      anonymous = true
    }
})

-- to create an anonymous context of a different kind
local deviceContext = ld.makeContext({
    device    = {
      key       = "example-device-key",
      anonymous = true
    }
})

```
