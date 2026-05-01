---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-alias-events-2-0-syntax-working-with-associated-contexts
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, working with associated contexts in section \"Understanding changes to alias events\""
---

```lua
-- create tables with the context attribute information
local device = {
  key = "example-device-key"
}

local user = {
  key = "example-user-key",
  attributes = {
    name = "Sandy"
  }
}

-- use them to create contexts at different points in your application

local deviceContext = ld.makeContext({
  device = device
})

client:identify(deviceContext)

local multiContext = ld.makeContext({
  device = device,
  user = user
})

client:identify(multiContext)
```
