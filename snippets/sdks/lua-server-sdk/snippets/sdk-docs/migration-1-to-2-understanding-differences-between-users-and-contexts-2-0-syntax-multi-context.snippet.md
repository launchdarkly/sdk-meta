---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-multi-context
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, multi-context in section \"Understanding differences between users and contexts\""
---

```lua
local context = ld.makeContext({
       user = {
           key = "example-user-key"
       },
       org = {
           key = "example-organization-key"
       }
})
```
