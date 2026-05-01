---
id: lua-server-sdk/sdk-docs/migration-1-to-2-understanding-changes-to-private-attributes-2-0-syntax-two-attributes-marked-private
sdk: lua-server-sdk
kind: reference
lang: lua
description: "2.0 syntax, two attributes marked private in section \"Understanding changes to private attributes\""
---

```lua
local configSomePrivate = {
    events = {
        privateAttributes = { "email", "address" }
    }
}
```
