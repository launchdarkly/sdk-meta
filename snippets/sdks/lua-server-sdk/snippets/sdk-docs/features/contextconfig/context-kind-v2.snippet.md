---
id: lua-server-sdk/sdk-docs/features/contextconfig/context-kind-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Context with a non-user kind for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
-- using makeContext to create a different kind of context
local orgContext = ld.makeContext({
    organization = {
      key = "example-organization-key",
      name = "Global Health Services"
    }
})
```
