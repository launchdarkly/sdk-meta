---
id: lua-server-sdk/sdk-docs/features/contextconfig/context-example-v2
sdk: lua-server-sdk
kind: reference
lang: lua
description: Context examples for Lua SDK v2.
validation:
  scaffold: lua-server-sdk/scaffolds/lua-syntax-only

---

```lua
-- using makeContext
local user1 = ld.makeContext({
    user = {
        key = "example-user-key-1",
        attributes = {
            firstName = "Sandy",
            lastName  = "Smith",
            email     = "sandy@example.com",
            groups    = { "Acme", "Global Health Services" }
        }
    }
})

-- using makeUser, which is deprecated,
-- to create an identical context (with unique key)
local user2 = ld.makeUser({
    key       = "example-user-key-2",
    firstName = "Sandy",
    lastName  = "Smith",
    email     = "sandy@example.com",
    custom    = {
        groups = { "Acme", "Global Health Services" }
    }
})

-- using makeContext to create a different kind of context
local orgContext = ld.makeContext({
    organization = {
      key = "example-organization-key",
      name = "Global Health Services"
    }
})
```
