---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-attribute-marked-private-for-one-context
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, attribute marked private for one context in section \"Understanding changes to private attributes\""
---

```lua
local user = ld.makeContext({
    user = {
      key       = "example-user-key",
      attributes = {
        firstName = "Sandy",
        lastName  = "Smith",
        email     = "sandy@example.com",
      },
      privateAttributes = { "email" }
    }
})
```
