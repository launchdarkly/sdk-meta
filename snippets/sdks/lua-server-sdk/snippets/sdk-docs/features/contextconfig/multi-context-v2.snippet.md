---
id: lua-server-sdk/sdk-docs/features/contextconfig/multi-context-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Multi-context example for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
-- using makeContext to create a multi-context
local context = ld.makeContext({
    user = {
      key = "example-user-key"
    },
    org = {
      key = "example-organization-key"
    }
})
```
