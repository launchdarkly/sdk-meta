---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-context-with-key
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, context with key in section \"Understanding differences between users and contexts\""
---

```lua
-- using makeUser, which is deprecated
local user1 = ld.makeUser({
    key       = "example-context-key1"
})

-- using makeContext to create an identical context (with unique key)
local user2 = ld.makeContext({
    user = {
      key       = "example-context-key2"
    }
})
```
