---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-differences-between-users-and-contexts-2-0-syntax-single-context-with-key
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, single context with key in section \"Understanding differences between users and contexts\""
---

```lua
local organization = ld.makeContext({
    organization = {
      key       = "example-organization-key"
    }
})
```
